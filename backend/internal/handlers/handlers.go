package handlers

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oli/birrwise-backend/internal/models"
)

func newID() string { b := make([]byte, 16); _, _ = rand.Read(b); return hex.EncodeToString(b) }
func now() string   { return time.Now().UTC().Format(time.RFC3339) }

func Register(r *gin.Engine, db *sql.DB) {
	a := r.Group("/api")
	a.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok", "database": "sqlite"}) })
	a.GET("/transactions", listTransactions(db))
	a.POST("/transactions", createTransaction(db))
	a.PUT("/transactions/:id", updateTransaction(db))
	a.DELETE("/transactions/:id", deleteTransaction(db))
	a.GET("/categories", listCategories(db))
	a.POST("/categories", createCategory(db))
	a.PUT("/categories/:id", updateCategory(db))
	a.DELETE("/categories/:id", deleteCategory(db))
	a.GET("/budgets", listBudgets(db))
	a.POST("/budgets", createBudget(db))
	a.PUT("/budgets/:id", updateBudget(db))
	a.DELETE("/budgets/:id", deleteBudget(db))
	a.GET("/preferences", getPreferences(db))
	a.PUT("/preferences", updatePreferences(db))
}

func listTransactions(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT id,title,amount,type,category_id,date,payment_method,description,created_at,updated_at FROM transactions ORDER BY date DESC,created_at DESC")
		if err != nil {
			serverError(c, err)
			return
		}
		defer rows.Close()
		items := []models.Transaction{}
		for rows.Next() {
			var x models.Transaction
			if err := rows.Scan(&x.ID, &x.Title, &x.Amount, &x.Type, &x.CategoryID, &x.Date, &x.PaymentMethod, &x.Description, &x.CreatedAt, &x.UpdatedAt); err != nil {
				serverError(c, err)
				return
			}
			items = append(items, x)
		}
		c.JSON(http.StatusOK, items)
	}
}
func createTransaction(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var x models.Transaction
		if err := c.ShouldBindJSON(&x); err != nil {
			badRequest(c, "Invalid transaction")
			return
		}
		if err := validateTransaction(db, x); err != nil {
			badRequest(c, err.Error())
			return
		}
		x.ID = newID()
		x.CreatedAt = now()
		x.UpdatedAt = x.CreatedAt
		_, err := db.Exec("INSERT INTO transactions VALUES(?,?,?,?,?,?,?,?,?,?)", x.ID, x.Title, x.Amount, x.Type, x.CategoryID, x.Date, x.PaymentMethod, x.Description, x.CreatedAt, x.UpdatedAt)
		if err != nil {
			serverError(c, err)
			return
		}
		c.JSON(http.StatusCreated, x)
	}
}
func updateTransaction(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var x models.Transaction
		if err := c.ShouldBindJSON(&x); err != nil {
			badRequest(c, "Invalid transaction")
			return
		}
		if err := validateTransaction(db, x); err != nil {
			badRequest(c, err.Error())
			return
		}
		x.ID = c.Param("id")
		x.UpdatedAt = now()
		res, err := db.Exec("UPDATE transactions SET title=?,amount=?,type=?,category_id=?,date=?,payment_method=?,description=?,updated_at=? WHERE id=?", x.Title, x.Amount, x.Type, x.CategoryID, x.Date, x.PaymentMethod, x.Description, x.UpdatedAt, x.ID)
		if err != nil {
			serverError(c, err)
			return
		}
		if n, _ := res.RowsAffected(); n == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
			return
		}
		c.JSON(http.StatusOK, x)
	}
}
func deleteTransaction(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := db.Exec("DELETE FROM transactions WHERE id=?", c.Param("id"))
		if err != nil {
			serverError(c, err)
			return
		}
		if n, _ := res.RowsAffected(); n == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
			return
		}
		c.Status(http.StatusNoContent)
	}
}
func validateTransaction(db *sql.DB, x models.Transaction) error {
	if x.Title == "" {
		return &validationError{"Title is required"}
	}
	if x.Amount <= 0 {
		return &validationError{"Amount must be greater than zero"}
	}
	if x.Type != "income" && x.Type != "expense" {
		return &validationError{"Type must be income or expense"}
	}
	if x.Date == "" {
		return &validationError{"Date is required"}
	}
	var count int
	if err := db.QueryRow("SELECT COUNT(*) FROM categories WHERE id=? AND type=?", x.CategoryID, x.Type).Scan(&count); err != nil {
		return err
	}
	if count == 0 {
		return &validationError{"Category is invalid for this transaction type"}
	}
	return nil
}

type validationError struct{ message string }

func (e *validationError) Error() string { return e.message }
func badRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{"error": message})
}
func serverError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
}

func listCategories(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT id,name,type,color,icon,is_default FROM categories ORDER BY name")
		if err != nil {
			serverError(c, err)
			return
		}
		defer rows.Close()
		items := []models.Category{}
		for rows.Next() {
			var x models.Category
			var defaultFlag int
			if err := rows.Scan(&x.ID, &x.Name, &x.Type, &x.Color, &x.Icon, &defaultFlag); err != nil {
				serverError(c, err)
				return
			}
			x.IsDefault = defaultFlag == 1
			items = append(items, x)
		}
		c.JSON(http.StatusOK, items)
	}
}
func createCategory(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var x models.Category
		if err := c.ShouldBindJSON(&x); err != nil || x.Name == "" || (x.Type != "income" && x.Type != "expense") {
			badRequest(c, "Name and valid type are required")
			return
		}
		x.ID = newID()
		if x.Color == "" {
			x.Color = "#64748b"
		}
		if x.Icon == "" {
			x.Icon = "Ellipsis"
		}
		_, err := db.Exec("INSERT INTO categories(id,name,type,color,icon,is_default) VALUES(?,?,?,?,?,0)", x.ID, x.Name, x.Type, x.Color, x.Icon)
		if err != nil {
			serverError(c, err)
			return
		}
		c.JSON(http.StatusCreated, x)
	}
}
func updateCategory(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var x models.Category
		if err := c.ShouldBindJSON(&x); err != nil || x.Name == "" || (x.Type != "income" && x.Type != "expense") {
			badRequest(c, "Name and valid type are required")
			return
		}
		res, err := db.Exec("UPDATE categories SET name=?,type=?,color=?,icon=? WHERE id=?", x.Name, x.Type, x.Color, x.Icon, c.Param("id"))
		if err != nil {
			serverError(c, err)
			return
		}
		if n, _ := res.RowsAffected(); n == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
			return
		}
		x.ID = c.Param("id")
		c.JSON(http.StatusOK, x)
	}
}
func deleteCategory(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var inUse int
		db.QueryRow("SELECT COUNT(*) FROM transactions WHERE category_id=?", c.Param("id")).Scan(&inUse)
		if inUse > 0 {
			badRequest(c, "Category is in use and cannot be deleted")
			return
		}
		res, err := db.Exec("DELETE FROM categories WHERE id=?", c.Param("id"))
		if err != nil {
			serverError(c, err)
			return
		}
		if n, _ := res.RowsAffected(); n == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
			return
		}
		c.Status(http.StatusNoContent)
	}
}

func listBudgets(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT id,category_id,month,limit_amount FROM budgets ORDER BY month DESC")
		if err != nil {
			serverError(c, err)
			return
		}
		defer rows.Close()
		items := []models.Budget{}
		for rows.Next() {
			var x models.Budget
			if err := rows.Scan(&x.ID, &x.CategoryID, &x.Month, &x.LimitAmount); err != nil {
				serverError(c, err)
				return
			}
			items = append(items, x)
		}
		c.JSON(http.StatusOK, items)
	}
}
func createBudget(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var x models.Budget
		if err := c.ShouldBindJSON(&x); err != nil || x.CategoryID == "" || x.Month == "" || x.LimitAmount <= 0 {
			badRequest(c, "Category, month, and valid limit are required")
			return
		}
		x.ID = newID()
		_, err := db.Exec("INSERT INTO budgets VALUES(?,?,?,?)", x.ID, x.CategoryID, x.Month, x.LimitAmount)
		if err != nil {
			if err.Error() == "constraint failed: UNIQUE constraint failed: budgets.category_id, budgets.month" {
				badRequest(c, "A budget already exists for this category and month")
				return
			}
			serverError(c, err)
			return
		}
		c.JSON(http.StatusCreated, x)
	}
}
func updateBudget(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var x models.Budget
		if err := c.ShouldBindJSON(&x); err != nil || x.CategoryID == "" || x.Month == "" || x.LimitAmount <= 0 {
			badRequest(c, "Category, month, and valid limit are required")
			return
		}
		res, err := db.Exec("UPDATE budgets SET category_id=?,month=?,limit_amount=? WHERE id=?", x.CategoryID, x.Month, x.LimitAmount, c.Param("id"))
		if err != nil {
			serverError(c, err)
			return
		}
		if n, _ := res.RowsAffected(); n == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Budget not found"})
			return
		}
		x.ID = c.Param("id")
		c.JSON(http.StatusOK, x)
	}
}
func deleteBudget(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := db.Exec("DELETE FROM budgets WHERE id=?", c.Param("id"))
		if err != nil {
			serverError(c, err)
			return
		}
		if n, _ := res.RowsAffected(); n == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Budget not found"})
			return
		}
		c.Status(http.StatusNoContent)
	}
}

func getPreferences(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var x models.Preferences
		if err := db.QueryRow("SELECT currency,language,theme FROM preferences WHERE id=1").Scan(&x.Currency, &x.Language, &x.Theme); err != nil {
			serverError(c, err)
			return
		}
		c.JSON(http.StatusOK, x)
	}
}
func updatePreferences(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var x models.Preferences
		if err := c.ShouldBindJSON(&x); err != nil {
			badRequest(c, "Invalid preferences")
			return
		}
		if x.Currency == "" {
			x.Currency = "ETB"
		}
		if x.Language != "en" && x.Language != "am" {
			x.Language = "en"
		}
		if x.Theme != "light" && x.Theme != "dark" {
			x.Theme = "light"
		}
		_, err := db.Exec("UPDATE preferences SET currency=?,language=?,theme=? WHERE id=1", x.Currency, x.Language, x.Theme)
		if err != nil {
			serverError(c, err)
			return
		}
		c.JSON(http.StatusOK, x)
	}
}

package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"

	_ "modernc.org/sqlite"
)

func Open(path string) (*sql.DB, error) {
	if path == "" {
		path = "birrwise.db"
	}
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return nil, err
	}
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(1)
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, migrate(db)
}

func migrate(db *sql.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS categories (id TEXT PRIMARY KEY, name TEXT NOT NULL, type TEXT NOT NULL CHECK(type IN ('income','expense')), color TEXT NOT NULL, icon TEXT NOT NULL DEFAULT 'Ellipsis', is_default INTEGER NOT NULL DEFAULT 0);
	CREATE TABLE IF NOT EXISTS transactions (id TEXT PRIMARY KEY, title TEXT NOT NULL, amount REAL NOT NULL CHECK(amount > 0), type TEXT NOT NULL CHECK(type IN ('income','expense')), category_id TEXT NOT NULL REFERENCES categories(id), date TEXT NOT NULL, payment_method TEXT NOT NULL, description TEXT NOT NULL DEFAULT '', created_at TEXT NOT NULL, updated_at TEXT NOT NULL);
	CREATE TABLE IF NOT EXISTS budgets (id TEXT PRIMARY KEY, category_id TEXT NOT NULL REFERENCES categories(id), month TEXT NOT NULL, limit_amount REAL NOT NULL CHECK(limit_amount > 0), UNIQUE(category_id, month));
	CREATE TABLE IF NOT EXISTS preferences (id INTEGER PRIMARY KEY CHECK(id = 1), currency TEXT NOT NULL DEFAULT 'ETB', language TEXT NOT NULL DEFAULT 'en', theme TEXT NOT NULL DEFAULT 'light');
	CREATE INDEX IF NOT EXISTS idx_transactions_date ON transactions(date);
	CREATE INDEX IF NOT EXISTS idx_transactions_type_date ON transactions(type, date);
	`
	if _, err := db.Exec(schema); err != nil {
		return fmt.Errorf("create schema: %w", err)
	}
	seedCategories(db)
	if err := seedData(db); err != nil {
		return fmt.Errorf("seed demo data: %w", err)
	}
	_, err := db.Exec("INSERT OR IGNORE INTO preferences(id,currency,language,theme) VALUES(1,'ETB','en','light')")
	return err
}

func seedCategories(db *sql.DB) {
	rows := []struct{ id, name, kind, color, icon string }{
		{"food", "Food and groceries", "expense", "#f59e0b", "Utensils"}, {"transport", "Transportation", "expense", "#3b82f6", "Car"}, {"housing", "Housing and rent", "expense", "#8b5cf6", "House"}, {"utilities", "Electricity and water", "expense", "#06b6d4", "Zap"}, {"internet", "Internet and airtime", "expense", "#ec4899", "Wifi"}, {"health", "Healthcare", "expense", "#ef4444", "HeartPulse"}, {"education", "Education", "expense", "#6366f1", "GraduationCap"}, {"clothing", "Clothing", "expense", "#14b8a6", "Shirt"}, {"entertainment", "Entertainment", "expense", "#a855f7", "Ticket"}, {"family", "Family expenses", "expense", "#f97316", "Users"}, {"other-expense", "Other", "expense", "#64748b", "Ellipsis"}, {"salary", "Salary", "income", "#10b981", "Briefcase"}, {"business", "Business income", "income", "#0ea5e9", "Store"}, {"freelance", "Freelancing", "income", "#22c55e", "Laptop"}, {"other-income", "Other income", "income", "#64748b", "WalletCards"},
	}
	for _, r := range rows {
		_, _ = db.Exec("INSERT OR IGNORE INTO categories(id,name,type,color,icon,is_default) VALUES(?,?,?,?,?,1)", r.id, r.name, r.kind, r.color, r.icon)
	}
}

func seedData(db *sql.DB) error {
	now := time.Now()
	timestamp := now.UTC().Format(time.RFC3339)
	currentMonth := now.Format("2006-01")
	date := func(daysAgo int) string {
		return now.AddDate(0, 0, -daysAgo).Format("2006-01-02")
	}

	transactions := []struct {
		id, title, kind, category, paymentMethod, description, date string
		amount                                                      float64
	}{
		{"seed-monthly-salary", "Monthly salary", "income", "salary", "Bank transfer", "November salary payment", date(2), 28000},
		{"seed-freelance-project", "Website design project", "income", "freelance", "CBE Birr", "Freelance client payment", date(8), 6500},
		{"seed-business-sales", "Weekend market sales", "income", "business", "Telebirr", "Small business revenue", date(12), 3200},
		{"seed-rent", "Apartment rent", "expense", "housing", "Bank transfer", "Monthly apartment rent", date(18), 8500},
		{"seed-groceries", "Mercato groceries", "expense", "food", "Telebirr", "Vegetables, grains, and household food", date(3), 1850},
		{"seed-taxi", "Taxi fares", "expense", "transport", "Cash", "Daily transport for the week", date(5), 720},
		{"seed-electricity", "Electricity and water", "expense", "utilities", "CBE Birr", "Monthly utility payment", date(9), 980},
		{"seed-internet", "Internet and airtime", "expense", "internet", "Telebirr", "Home internet and mobile airtime", date(7), 1250},
		{"seed-family", "Family support", "expense", "family", "Cash", "Support for family expenses", date(14), 1500},
		{"seed-entertainment", "Cinema and coffee", "expense", "entertainment", "Telebirr", "Weekend leisure", date(6), 650},
	}

	for _, transaction := range transactions {
		_, err := db.Exec(`INSERT OR IGNORE INTO transactions
			(id,title,amount,type,category_id,date,payment_method,description,created_at,updated_at)
			VALUES(?,?,?,?,?,?,?,?,?,?)`,
			transaction.id, transaction.title, transaction.amount, transaction.kind,
			transaction.category, transaction.date, transaction.paymentMethod,
			transaction.description, timestamp, timestamp)
		if err != nil {
			return err
		}
	}

	budgets := []struct {
		id, category string
		amount       float64
	}{
		{"budget-food-current", "food", 3500},
		{"budget-transport-current", "transport", 2200},
		{"budget-housing-current", "housing", 9000},
		{"budget-internet-current", "internet", 1800},
		{"budget-entertainment-current", "entertainment", 1500},
	}

	for _, budget := range budgets {
		if _, err := db.Exec("INSERT OR IGNORE INTO budgets(id,category_id,month,limit_amount) VALUES(?,?,?,?)", budget.id+"-"+currentMonth, budget.category, currentMonth, budget.amount); err != nil {
			return err
		}
	}

	return nil
}

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

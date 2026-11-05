package models

type Transaction struct {
	ID            string  `json:"id" db:"id"`
	Title         string  `json:"title" db:"title"`
	Amount        float64 `json:"amount" db:"amount"`
	Type          string  `json:"type" db:"type"`
	CategoryID    string  `json:"categoryId" db:"category_id"`
	Date          string  `json:"date" db:"date"`
	PaymentMethod string  `json:"paymentMethod" db:"payment_method"`
	Description   string  `json:"description" db:"description"`
	CreatedAt     string  `json:"createdAt" db:"created_at"`
	UpdatedAt     string  `json:"updatedAt" db:"updated_at"`
}

type Category struct {
	ID        string `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Type      string `json:"type" db:"type"`
	Color     string `json:"color" db:"color"`
	Icon      string `json:"icon" db:"icon"`
	IsDefault bool   `json:"isDefault" db:"is_default"`
}

type Budget struct {
	ID          string  `json:"id" db:"id"`
	CategoryID  string  `json:"categoryId" db:"category_id"`
	Month       string  `json:"month" db:"month"`
	LimitAmount float64 `json:"limitAmount" db:"limit_amount"`
}

type Preferences struct {
	Currency string `json:"currency"`
	Language string `json:"language"`
	Theme    string `json:"theme"`
}

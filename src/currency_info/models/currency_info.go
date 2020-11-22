package models

type CurrInfo struct {
	CurrencyName string `json:"currencyName" db:"name"`
	StockName string `json:"stockName" db:"stock_name"`
	Description string `json:"description" db:"description"`
}

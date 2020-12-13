package models

type MinCurrencyInfo struct {
	CmcId int `json:"id"`
	Name string `json:"name"`
	Symbol string `json:"symbol"`
	Rank int `json:"rank"`
	Logo string `json:"logo"`
	Category string `json:"category"`
	LastUpdated string `json:"last_updated" db:"last_updated"`
	PercentChange1h float64 `json:"percent_change_1h" db:"percent_change_1h"`
	PercentChange24h float64 `json:"percent_change_24h" db:"percent_change_24h"`
	PercentChange7d float64 `json:"percent_change_7d" db:"percent_change_7d"`
	Cost []Fiat `json:"cost"`
}

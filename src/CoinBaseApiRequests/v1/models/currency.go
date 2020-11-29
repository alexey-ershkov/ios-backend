package models

type Fiat struct {
	CmcId int `json:"cmc_id" db:"cmc_id"`
	Symbol string `json:"symbol" db:"fiat_symbol"`
	Price float64 `json:"price" db:"price"`
	Volume24h float64 `json:"volume_24h" db:"volume"`
	LastUpdated string `json:"last_updated" db:"last_updated"`
}

type CurrCryptoInfo struct {
	CmcId int `json:"id" db:"cmc_id"`
	MaxSupply float64 `json:"max_supply" db:"max"`
	CirculatingSupply float64 `json:"circulating_supply" db:"in_market"`
	TotalSupply float64 `json:"total_supply" db:"mined"`
	LastUpdated string `json:"last_updated" db:"last_updated"`
	PercentChange1h float64 `json:"percent_change_1h" db:"percent_change_1h"`
	PercentChange24h float64 `json:"percent_change_24h" db:"percent_change_24h"`
	PercentChange7d float64 `json:"percent_change_7d" db:"percent_change_7d"`
	CostInFiats []Fiat
}

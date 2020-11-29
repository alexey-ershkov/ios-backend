package models

type FiatModel struct {
	FiatId int `json:"id" db:"cmc_fiat_id"`
	Name string `json:"name" db:"name"`
	Sign string `json:"sign" db:"sign"`
	Symbol string `json:"symbol" db:"symbol"`
}

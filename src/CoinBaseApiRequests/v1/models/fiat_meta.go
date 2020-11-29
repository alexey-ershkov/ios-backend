package models

type FiatMeta struct {
	Name string `json:"name" db:"name"`
	Sign string `json:"sign" db:"sign"`
	Symbol string `json:"symbol" db:"symbol"`
}

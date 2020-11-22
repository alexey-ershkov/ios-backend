package models

type Currency struct {
	CurrencyName string `json:"currencyName"`
	StockName string `json:"stockName"`
	Cost float32 `json:"cost"`
	DidGrow bool `json:"didGrow"`
	ChangeValueInPercents float32 `json:"changeValueInPercents"`
	Description string `json:"description"`
	ConvertionCurrencyName string `json:"convertionCurrencyName"`
}

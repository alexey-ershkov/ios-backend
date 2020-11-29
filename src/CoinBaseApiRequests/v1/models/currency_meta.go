package models

type PlatformMeta struct {
	CmcId int `json:"id"`
	Name string `json:"name"`
	Symbol string `json:"symbol"`
	TokenAddress string `json:"token_address"`
}

type Urls struct {
	Website []string `json:"website"`
	TechnicalDoc []string `json:"technical_doc"`
	Twitter []string `json:"twitter"`
	Reddit []string `json:"reddit"`
	SourceCode []string `json:"source_code"`
	
}

type CurrencyMeta struct {
	CmcId int `json:"id" db:"cmc_id"`
	Name string `json:"name" db:"name"`
	Symbol string `json:"symbol" db:"symbol"`
	Rank int `json:"rank" db:"rank"`
	Logo string `json:"logo" db:"logo"`
	Category string `json:"category" db:"category"`
	Description string `json:"description" db:"description"`
	DateAdded string `json:"date_added" db:"date_added"`
	Platform PlatformMeta `json:"platform" db:"platform"`
	Urls Urls `json:"urls"`
}

type CurrencyMetaDB struct {
	CmcId int `json:"id" db:"cmc_id"`
	Name string `json:"name" db:"name"`
	Symbol string `json:"symbol" db:"symbol"`
	Rank int `json:"rank" db:"rank"`
	Logo string `json:"logo" db:"logo"`
	Category string `json:"category" db:"category"`
	Description string `json:"description" db:"description"`
	DateAdded string `json:"date_added" db:"date_added"`
	PlatformCmcId int `json:"platform_cmc_id" db:"platform_cmc_id"`
	PlatformSymbol string `json:"platform_symbol" db:"platform_symbol"`
	PlatformName string `json:"platform_name" db:"platform_name"`
	PlatformTokenAddress string `json:"platform_token_address" db:"platform_token_address"`
	Website string `json:"website" db:"website"`
	Doc string `json:"doc" db:"doc"`
	Twitter string `json:"twitter" db:"twitter"`
	Reddit string `json:"reddit" db:"reddit"`
	SourceCode string `json:"source_code" db:"source_code"`
}

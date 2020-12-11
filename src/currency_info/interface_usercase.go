package currency_info

import "ios-backend/src/currency_info/models"

type CurrUCase interface {
	GetCurrencyByName(name string) (*models.Currency, error)
	GetCurrencyListByStockNames(names []string)([]models.MinCurrencyInfo, error)
}

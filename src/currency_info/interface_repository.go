package currency_info

import "ios-backend/src/currency_info/models"

type CurrRepo interface {
	GetCurrInfoByNameOrStockName(name string) (*models.CurrInfo, error)
}


package currency_info

import "ios-backend/src/CoinBaseApiRequests/v1/models"
import im "ios-backend/src/currency_info/models"

type CurrRepo interface {
	GetMetaInfoByNameStockNameOrCmcId(name string) (*models.CurrencyMetaDB, error)
	GetCurrInfoByCmcId(id int) (*im.CurrCryptoInfo, error)
	GetCurrInfoInFiatByCmcId(id int) ([]im.Fiat, error)
	GetCurrencyListByStockNameOrName(names []string)([]im.MinCurrencyInfo, error)
}


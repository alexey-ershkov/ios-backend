package usecase

import (
	"ios-backend/src/currency_info"
	"ios-backend/src/currency_info/models"
)

type CurrUsecase struct {
	Repo      currency_info.CurrRepo
}

func NewCurrUsecase(currRepo currency_info.CurrRepo) currency_info.CurrUCase {
	return CurrUsecase{Repo: currRepo}
}

func (uc CurrUsecase) GetCurrencyByName(name string) (*models.Currency, error) {
	curr := &models.Currency{}
	info, err := uc.Repo.GetCurrInfoByNameOrStockName(name)
	if err != nil {
		return nil, err
	}

	curr.CurrencyName = info.CurrencyName
	curr.StockName = info.StockName
	curr.Description = info.Description
	curr.ChangeValueInPercents = 3.61
	curr.Cost = 23.4
	curr.DidGrow = true
	curr.ConvertionCurrencyName = "USD"
	return curr, nil
}
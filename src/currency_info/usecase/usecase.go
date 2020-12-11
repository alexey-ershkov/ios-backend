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
	meta, err := uc.Repo.GetMetaInfoByNameStockNameOrCmcId(name)
	if err != nil {
		return nil, err
	}
	curr.CmcId = meta.CmcId
	curr.Name = meta.Name
	curr.Symbol = meta.Symbol
	curr.Rank = meta.Rank
	curr.Logo = meta.Logo
	curr.DateAdded = meta.DateAdded
	curr.Category = meta.Category
	curr.Description = meta.Description
	pl := models.PlatformMeta{
		CmcId:        meta.PlatformCmcId,
		Name:         meta.PlatformName,
		Symbol:       meta.PlatformSymbol,
		TokenAddress: meta.PlatformTokenAddress,
	}
	curr.Platform = pl
	urls := models.Urls{
		Website:      meta.Website,
		TechnicalDoc: meta.Doc,
		Twitter:      meta.Twitter,
		Reddit:       meta.Reddit,
		SourceCode:   meta.SourceCode,
	}
	curr.Urls = urls
	currInfo, err := uc.Repo.GetCurrInfoByCmcId(curr.CmcId)
	if err != nil {
		return nil, err
	}

	fiatInfo, err := uc.Repo.GetCurrInfoInFiatByCmcId(curr.CmcId)
	if err != nil {
		return nil, err
	}

	currInfo.CostInFiats = fiatInfo
	curr.CurrCryptoInfo = *currInfo

	return curr, nil
}

func (uc CurrUsecase) GetCurrencyListByStockNames(names []string)([]models.MinCurrencyInfo, error) {
	currInfo, err := uc.Repo.GetCurrencyListByStockNameOrName(names)
	if err != nil {
		return nil, err
	}

	return currInfo, nil
}
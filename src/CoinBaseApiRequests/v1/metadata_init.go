package v1

import (
	"github.com/jmoiron/sqlx"
	"ios-backend/src/CoinBaseApiRequests/v1/models"
)

func convertArray(data []string) string {
	if len(data) > 0 {
		return data[0]
	}
	return ""
}

func convertToDBFormat(meta models.CurrencyMeta) models.CurrencyMetaDB {
	return models.CurrencyMetaDB{
		CmcId:                meta.CmcId,
		Name:                 meta.Name,
		Symbol:               meta.Symbol,
		Rank:                 meta.Rank,
		Logo:                 meta.Logo,
		Category:             meta.Category,
		Description:          meta.Description,
		DateAdded:            meta.DateAdded,
		PlatformCmcId:        meta.Platform.CmcId,
		PlatformSymbol:       meta.Platform.Symbol,
		PlatformName:         meta.Platform.Name,
		PlatformTokenAddress: meta.Platform.TokenAddress,
		Website:              convertArray(meta.Urls.Website),
		Doc:                  convertArray(meta.Urls.TechnicalDoc),
		Twitter:              convertArray(meta.Urls.Twitter),
		Reddit:               convertArray(meta.Urls.Reddit),
		SourceCode:           convertArray(meta.Urls.SourceCode),
	}
}

func insertMetadataInDb(conn *sqlx.DB, data models.CurrencyMetaDB) error {

	query := "insert into currency_info (name, symbol, cmc_id, rank, logo, date_added, category, description, " +
		"platform_cmc_id, platform_name, platform_symbol, platform_token_address, website, doc, twitter, reddit, " +
		"source_code)" +
		"values " +
		"(:name, :symbol, :cmc_id, :rank, :logo, :date_added, :category, :description, :platform_cmc_id," +
		" :platform_name, :platform_symbol, :platform_token_address, :website, :doc, :twitter, :reddit," +
		":source_code); "

	_, err := conn.NamedExec(query, data)
	if err != nil {
		return err
	}

	return nil
}

func GetCurrencyMetadata(conn *sqlx.DB) error {
	api, err := NewCurrencyApi()
	if err != nil {
		return err
	}

	data, err := api.GetMetadata()
	if err != nil {
		return err
	}

	for _, singleData := range data {
		singleDataDbReady := convertToDBFormat(singleData)
		err := insertMetadataInDb(conn, singleDataDbReady)
		if err != nil {
			return err
		}
	}

	return nil
}

package v1

import (
	"github.com/jmoiron/sqlx"
	"ios-backend/src/CoinBaseApiRequests/v1/models"
)

func insertInDB(conn *sqlx.DB, model models.CurrCryptoInfo) error {

	currCryptoInsertQuery := "insert into curr_crypto_info " +
		"(cmc_id, max, in_market, mined, last_updated, percent_change_1h, percent_change_24h, percent_change_7d) " +
		"values " +
		"(:cmc_id, :max, :in_market, :mined, :last_updated, " +
		":percent_change_1h, :percent_change_24h, :percent_change_7d) " +
		"on conflict (cmc_id) do update set " +
		"cmc_id = excluded.cmc_id,max = excluded.max, " +
		"in_market = excluded.in_market,mined = excluded.mined, " +
		"last_updated = excluded.last_updated, " +
		"percent_change_1h = excluded.percent_change_1h, " +
		"percent_change_24h = excluded.percent_change_24h, " +
		"percent_change_7d = excluded.percent_change_7d;"

	currFiatsInsertQuery := "insert into curr_crypto_info_in_fiat " +
		"(fiat_symbol, cmc_id, price, volume, last_updated) " +
		"values " +
		"(:fiat_symbol, :cmc_id, :price, :volume, :last_updated) " +
		"on conflict (fiat_symbol, cmc_id) do update " +
		"set price = excluded.price, volume = excluded.volume, last_updated = excluded.last_updated;"

	_, err := conn.NamedExec(currCryptoInsertQuery, model)
	if err != nil {
		return err
	}

	for _, fiat := range model.CostInFiats {
		_, err := conn.NamedExec(currFiatsInsertQuery, fiat)
		if err != nil {
			return err
		}
	}

	return nil
}

func UpdateCryptoInfo(conn *sqlx.DB) error {
	api, err := NewCurrencyApi()
	if err != nil {
		return err
	}

	data, err := api.GetCurrCurrencyInfo()
	if err != nil {
		return err
	}

	for _, elem := range data {
		err := insertInDB(conn, elem)
		if err != nil {
			return err
		}
	}

	return nil
}

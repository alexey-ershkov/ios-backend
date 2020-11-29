package v1

import (
	"github.com/jmoiron/sqlx"
	"ios-backend/src/CoinBaseApiRequests/v1/models"
)

func insertFiatMetadataInDb(conn *sqlx.DB, data models.FiatMeta) error {

	query := "insert into fiat_info (name, sign, symbol) " +
		"values (:name, :sign, :symbol);"

	_, err := conn.NamedExec(query, data)
	if err != nil {
		return err
	}

	return nil
}

func GetFiatMetadata(conn *sqlx.DB) error {

	api, err := NewCurrencyApi()
	if err != nil {
		return err
	}

	data, err := api.GetFiatMetadata()
	if err != nil {
		return err
	}


	for _, singleData := range data {
		err := insertFiatMetadataInDb(conn, singleData)
		if err != nil {
			return err
		}
	}

	return nil
}

package v1

import "database/sql"

func UpdateCryptoInfo(conn *sql.Conn) error {
	api, err := NewCurrencyApi()
	if err != nil {
		return err
	}

	_, err = api.GetFiatMetadata()
	if err != nil {
		return err
	}

	return nil
}

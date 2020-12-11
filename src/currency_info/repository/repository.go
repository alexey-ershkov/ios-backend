package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"ios-backend/src/CoinBaseApiRequests/v1/models"
	"ios-backend/src/currency_info"
	im "ios-backend/src/currency_info/models"
	"strings"
)

type CurrRepo struct {
	conn *sqlx.DB
}

func NewCurrRepo(conn *sqlx.DB) currency_info.CurrRepo {
	return CurrRepo{conn}
}

func (r CurrRepo) GetMetaInfoByNameStockNameOrCmcId(name string) (*models.CurrencyMetaDB, error) {
	info := &models.CurrencyMetaDB{}
	sqlQuery := "select * from currency_info where citext(cmc_id) = $1 or name = $1 or symbol = $1"

	err := r.conn.Get(info, sqlQuery, name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return info, nil
}

func (r CurrRepo) GetCurrInfoByCmcId(id int) (*im.CurrCryptoInfo, error) {
	info := &im.CurrCryptoInfo{}
	sqlQuery := "select * from curr_crypto_info cci where cci.cmc_id = $1;"

	err := r.conn.Get(info, sqlQuery, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return info, nil
}

func (r CurrRepo) GetCurrInfoInFiatByCmcId(id int) ([]im.Fiat, error) {
	info := make([]im.Fiat, 0)
	sqlQuery := "select * from curr_crypto_info_in_fiat cif " +
		"join fiat_info fi on fi.symbol = cif.fiat_symbol " +
		"where cif.cmc_id = $1 " +
		"order by fi.symbol ;"

	rows, err := r.conn.Query(sqlQuery, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for rows.Next() {
		buf := im.Fiat{}
		err = rows.Scan(
			&buf.Symbol,
			&buf.CmcId,
			&buf.Price,
			&buf.Volume24h,
			&buf.LastUpdated,
			&buf.Symbol,
			&buf.Sign,
			&buf.Name)
		if err != nil {
			return nil, err
		}
		info = append(info, buf)
	}

	return info, nil
}

func (r CurrRepo) GetCurrencyListByStockNameOrName(names []string)([]im.MinCurrencyInfo, error){
	info := make([]im.MinCurrencyInfo, 0)
	sqlQuery := "select ci.cmc_id, name, symbol, rank, logo, category, last_updated, " +
		"percent_change_1h, percent_change_24h, percent_change_7d from currency_info ci " +
		"join curr_crypto_info cci on ci.cmc_id = cci.cmc_id "

	if names != nil {
		n := strings.Join(names, "', '")
		sqlQuery += "where ci.symbol in ('" + n + "') or ci.name in ('" + n + "');"
	} else {
		sqlQuery += ";"
	}


	rows, err := r.conn.Query(sqlQuery)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		buf := im.MinCurrencyInfo{}
		err = rows.Scan(&buf.CmcId,
			&buf.Name,
			&buf.Symbol,
			&buf.Rank,
			&buf.Logo,
			&buf.Category,
			&buf.LastUpdated,
			&buf.PercentChange1h,
			&buf.PercentChange24h,
			&buf.PercentChange7d)

		if err != nil {
			return nil, err
		}

		fiat, err := r.GetCurrInfoInFiatByCmcId(buf.CmcId)
		if err != nil {
			return nil, err
		}

		buf.Cost = fiat
		info = append(info, buf)
	}

	return info, nil
}

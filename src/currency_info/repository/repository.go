package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"ios-backend/src/currency_info"
	"ios-backend/src/currency_info/models"
)

type CurrRepo struct {
	conn *sqlx.DB
}

func NewCurrRepo(conn *sqlx.DB) currency_info.CurrRepo {
	return CurrRepo{conn}
}

func (r CurrRepo) GetCurrInfoByNameOrStockName(name string) (*models.CurrInfo, error) {
	info := &models.CurrInfo{}
	sqlQuery := "select name, stock_name, description from currency_info where name = $1 or stock_name=$1;"

	err := r.conn.Get(info, sqlQuery, name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return info, nil
}

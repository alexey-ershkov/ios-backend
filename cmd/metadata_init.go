package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	cmc "ios-backend/src/CoinBaseApiRequests/v1"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable port=%s",
		"postgres", //docker,postgres
		"",         //docker, empty
		"postgres", //docker,postgres
		"5432") // для тестов на локалке

	if os.Getenv("IN_DOCKER") == "true" {
		connStr = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable port=%s",
			"docker", //docker,postgres
			"docker", //docker, empty
			"docker", //docker,postgres
			"5432")
	}

	conn, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Error().Msgf(err.Error())
		return
	}


	//err = cmc.GetCurrencyMetadata(conn)
	//if err != nil {
	//	log.Error().Msgf(err.Error())
	//}

	err = cmc.GetFiatMetadata(conn)
	if err != nil {
		log.Error().Msgf(err.Error())
	}
}

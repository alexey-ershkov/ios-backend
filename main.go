package main

import (
	"fmt"
	"github.com/joho/godotenv"
	v1 "ios-backend/src/CoinBaseApiRequests/v1"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"ios-backend/src/configs"
	"ios-backend/src/user/delivery"
	"ios-backend/src/user/repository"
	"ios-backend/src/user/usecase"

	currDelivery "ios-backend/src/currency_info/delivery"
	currRepo "ios-backend/src/currency_info/repository"
	currUCase "ios-backend/src/currency_info/usecase"
)

func updateCurr (conn *sqlx.DB) {
	err := v1.UpdateCryptoInfo(conn)
	if err != nil {
		log.Error().Msgf(err.Error())
	}
}

func updateCurrEvery(d time.Duration, conn *sqlx.DB) {
	for range time.Tick(d) {
		updateCurr(conn)
	}
}


func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	r := mux.NewRouter()
	timeoutContext := configs.Timeouts.ContextTimeout

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


	rep := repository.NewPostgresUserRepository(conn)
	ucase := usecase.NewUserUsecase(rep, timeoutContext)
	delivery.NewUserHandler(r, ucase)

	currR := currRepo.NewCurrRepo(conn)
	currUC := currUCase.NewCurrUsecase(currR)
	currDelivery.NewUserHandler(r, currUC)

	updateCurr(conn)
	go updateCurrEvery(configs.UPD_INTERVAL, conn)

	//static server
	r.PathPrefix(fmt.Sprintf("/%s/", configs.MEDIA_FOLDER)).Handler(
		http.StripPrefix(fmt.Sprintf("/%s/", configs.MEDIA_FOLDER),
			http.FileServer(http.Dir(configs.MEDIA_FOLDER))))

	http.Handle("/", r)
	srv := &http.Server{
		Handler:      r,
		Addr:         configs.SERVER_ADDRESS,
		WriteTimeout: configs.Timeouts.WriteTimeout,
		ReadTimeout:  configs.Timeouts.ReadTimeout,
	}
	fmt.Println("main server started at ", configs.SERVER_URL)
	log.Error().Msgf(srv.ListenAndServe().Error())
}

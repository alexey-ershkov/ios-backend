package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"ios-backend/src/configs"
	"ios-backend/src/user/delivery"
	"ios-backend/src/user/repository"
	"ios-backend/src/user/usecase"
)

func main() {
	r := mux.NewRouter()

	timeoutContext := configs.Timeouts.ContextTimeout

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable port=%s",
		"docker", //docker,postgres
		"docker", //docker, empty
		"docker", //docker,postgres
		"5432")
	conn, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Error().Msgf(err.Error())
		return
	}

	rep := repository.NewPostgresUserRepository(conn)
	ucase := usecase.NewUserUsecase(rep, timeoutContext)
	delivery.NewUserHandler(r, ucase)

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

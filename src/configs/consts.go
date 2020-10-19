package configs

import (
	"time"
)

const (
	SERVER_PORT    = ":5000"
	SERVER_URL     = "http://0.0.0.0:5000"
	MEDIA_FOLDER   = "media"
	SERVER_ADDRESS = "0.0.0.0:5000"
)

type timeouts struct {
	WriteTimeout   time.Duration
	ReadTimeout    time.Duration
	ContextTimeout time.Duration
}

var Timeouts = timeouts{
	WriteTimeout:   time.Second * 5,
	ReadTimeout:    time.Second * 5,
	ContextTimeout: time.Second * 5,
}

type postgresPreferencesStruct struct {
	User     string
	Password string
	Port     string
	Host     string
	DBName   string
}

var PostgresPreferences = postgresPreferencesStruct{
	User:     "postgres",
	Password: "postgres",
	Port:     "5432",
	Host:     "0.0.0.0",
	DBName:   "postgres",
}

//var PostgresPreferences = postgresPreferencesStruct{
//	User:     "postgres",
//	Password: "",
//	Port:     "5432",
//	Host:     "127.0.0.1",
//	DBName:   "postgres",
//}

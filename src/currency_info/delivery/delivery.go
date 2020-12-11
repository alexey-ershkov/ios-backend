package delivery

import (
	"github.com/gorilla/mux"
	"ios-backend/src/currency_info"
	"ios-backend/src/utills"
	"net/http"
	"strings"
)

type CurrencyHandler struct {
	UC     currency_info.CurrUCase
	ApiKey string
	ApiURL string
}

func NewUserHandler(r *mux.Router, UC currency_info.CurrUCase) {
	handler := CurrencyHandler{
		UC: UC,
	}
	r.HandleFunc("/api/currency/get", handler.GetCurrency).Methods(http.MethodGet)
	r.HandleFunc("/api/currency/list", handler.GetList).Methods(http.MethodGet)
}

func (ch CurrencyHandler) GetCurrency(w http.ResponseWriter, r *http.Request) {
	args, ok := r.URL.Query()["curr_name"]
	if !ok || len(args) < 1 {
		utills.SendServerError("No param in query", 404, w)
		return
	}
	name := args[0]

	currInfo, err := ch.UC.GetCurrencyByName(name)
	if err != nil {
		utills.SendServerError(err.Error(), 500, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	utills.SendOKAnswer(currInfo, w)
}

func (ch CurrencyHandler) GetList(w http.ResponseWriter, r *http.Request) {
	args, _ := r.URL.Query()["names"]

	if args != nil {
		args = strings.Split(args[0], ",")
	} else {
		args = nil
	}


	currInfo, err := ch.UC.GetCurrencyListByStockNames(args)
	if err != nil {
		utills.SendServerError(err.Error(), 500, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	utills.SendOKAnswer(currInfo, w)
}

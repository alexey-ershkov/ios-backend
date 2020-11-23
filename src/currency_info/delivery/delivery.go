package delivery

import (
	"github.com/gorilla/mux"
	"ios-backend/src/currency_info"
	"ios-backend/src/utills"
	"net/http"
)

type UserHandler struct {
	UC currency_info.CurrUCase
}

func NewUserHandler(r *mux.Router, UC currency_info.CurrUCase) {
	handler := UserHandler{
		UC: UC,
	}
	r.HandleFunc("/currency/get", handler.GetCurrency).Methods(http.MethodGet)
}

func (u UserHandler) GetCurrency (w http.ResponseWriter, r *http.Request) {
	args, ok := r.URL.Query()["curr_name"]
	if !ok || len(args) < 1 {
		utills.SendServerError("No param in query", 404, w)
		return
	}
	name := args[0]

	currInfo, err := u.UC.GetCurrencyByName(name)
	if err != nil {
		utills.SendServerError(err.Error(), 500, w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	utills.SendOKAnswer(currInfo, w)
}

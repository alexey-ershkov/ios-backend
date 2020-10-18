package utills

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

func SendOKAnswer(data interface{}, w http.ResponseWriter) {
	SendAnswerWithCode(data, 200, w)
}

func SendAnswerWithCode(data interface{}, code int, w http.ResponseWriter) {
	w.WriteHeader(code)
	serializedData, err := json.Marshal(data)
	if err != nil {
		log.Error().Msgf(err.Error())
		return
	}

	_, err = w.Write(serializedData)
	if err != nil {
		message := fmt.Sprintf("HttpResponse while writing is socket: %s", err.Error())
		log.Error().Msgf(message)
		return
	}
	log.Info().Msgf("Code message sent")
}

type ModelError struct {
	Message string `json:"message,omitempty"`
}

func SendServerError(errorMessage string, code int, w http.ResponseWriter) {
	log.Error().Msgf(errorMessage)
	w.WriteHeader(code)
	mes, _ := json.Marshal(ModelError{Message: errorMessage})
	w.Write(mes)
}

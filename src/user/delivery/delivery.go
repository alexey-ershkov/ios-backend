package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"ios-backend/src/configs"
	"ios-backend/src/user"
	"ios-backend/src/user/models"
	"ios-backend/src/utills"
)

type UserHandler struct {
	SUsecase user.Usecase
}

func NewUserHandler(r *mux.Router, us user.Usecase) {
	handler := UserHandler{
		SUsecase: us,
	}
	r.HandleFunc("/api/add", handler.AddUser).Methods(http.MethodPost)

}

func (s *UserHandler) fetchUser(r *http.Request) (models.User, error) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		return models.User{}, configs.ErrBadRequest
	}

	jsonData := r.FormValue("jsonData")
	if jsonData == "" || jsonData == "null" {
		return models.User{}, configs.ErrEmptyJSON
	}

	var UserObj models.User
	err = json.Unmarshal([]byte(jsonData), &UserObj)
	if err != nil {
		return models.User{}, configs.ErrBadJSON
	}
	if file, handler, err := r.FormFile("photo"); err == nil {
		filename, err := utills.SaveFile(file, handler, "Users")
		if err == nil {
			UserObj.Photo = fmt.Sprintf("%s/%s", configs.SERVER_URL, filename)
		}
	}

	return UserObj, nil
}

func (u UserHandler) AddUser(writer http.ResponseWriter, request *http.Request) {
	usr, err := u.fetchUser(request)
	if err != nil {
		utills.SendServerError(err.Error(), 500, writer)
		return
	}
	newUser, err := u.SUsecase.Add(request.Context(), usr)
	if err != nil {
		utills.SendServerError(err.Error(), 500, writer)
		return
	}

	tenYears := time.Now().Add(time.Hour * 24 * 30 * 100)
	endlessCookie := &http.Cookie{
		Name:       "user_id",
		Value:      strconv.Itoa(newUser.UserID),
		Path:       "/",
		RawExpires: "",
		Expires:    tenYears,
		HttpOnly:   true,
	}
	http.SetCookie(writer, endlessCookie)

	utills.SendOKAnswer(newUser, writer)
}

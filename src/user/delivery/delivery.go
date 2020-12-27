package delivery

import (
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
	r.HandleFunc("/api/user/add", handler.AddUser).Methods(http.MethodPost)
	r.HandleFunc("/api/user/get", handler.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/api/user/login", handler.LoginUser).Methods(http.MethodPost)
}

func (s *UserHandler) fetchUser(r *http.Request) (models.User, error) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		return models.User{}, configs.ErrBadRequest
	}

	UserObj := models.User{
		NickName: r.FormValue("Nickname"),
		Email:    r.FormValue("Email"),
		Password: r.FormValue("Password"),
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
	newUser.UserID = 0 // for security
	utills.SendOKAnswer(newUser, writer)
}

func (s *UserHandler) GetUser(writer http.ResponseWriter, request *http.Request) {
	userIdCookie, err := request.Cookie("user_id")
	if err != nil {
		utills.SendServerError(configs.ErrUserIsNotRegistered.Error(), http.StatusUnauthorized, writer)
		return
	}
	userId, err := strconv.Atoi(userIdCookie.Value)
	if err != nil {
		utills.SendServerError(configs.ErrUserIdIsNotNumber.Error(), http.StatusUnauthorized, writer)
		return
	}
	usr, err := s.SUsecase.GetCurrent(request.Context(), userId)
	if err != nil {
		utills.SendServerError(err.Error(), http.StatusUnauthorized, writer)
		return
	}
	utills.SendOKAnswer(usr, writer)
}

func (s *UserHandler) LoginUser(writer http.ResponseWriter, request *http.Request) {
	email, password, err := fetchEmailAndPassword(request)
	if err != nil {
		utills.SendServerError(err.Error(), http.StatusUnauthorized, writer)
		return
	}
	usr, err := s.SUsecase.GetUserByEmailAndPassword(request.Context(), email, password)
	if err != nil {
		utills.SendServerError(err.Error(), http.StatusUnauthorized, writer)
		return
	}
	tenYears := time.Now().Add(time.Hour * 24 * 30 * 100)
	endlessCookie := &http.Cookie{
		Name:       "user_id",
		Value:      strconv.Itoa(usr.UserID),
		Path:       "/",
		RawExpires: "",
		Expires:    tenYears,
		HttpOnly:   true,
	}
	http.SetCookie(writer, endlessCookie)
	utills.SendOKAnswer(usr, writer)
}

func fetchEmailAndPassword(r *http.Request) (string, string, error) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		return "", "", configs.ErrBadRequest
	}
	return r.FormValue("Email"), r.FormValue("Password"), nil
}

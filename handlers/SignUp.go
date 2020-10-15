package handlers

import (
	"backend/models"
	"errors"
	"fmt"
	"net/http"
)

func readSignUpRequest(r *http.Request) (models.UserValues, error) {
	user := models.UserValues{}
	var err error

	user.Login = r.URL.Query().Get("login")
	if len(user.Login) == 0 {
		err = errors.New("empty login")
		return user, err
	}

	user.Password = r.URL.Query().Get("password")
	if len(user.Password) == 0 {
		err = errors.New("empty password")
		return user, err
	}

	return user, nil
}

func addUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	userValues, err := readSignUpRequest(r)

	if err != nil {
		fmt.Println(err)
		return
	}

	user := models.User{
		Login: userValues.Login,
		Password: userValues.Password,
	}
	db.create
}
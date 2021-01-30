package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	username string `json: username`
	age      int    `json: age`
	email    string `json: email`
	password string `json: password`
	ID       string `json: userID`
}

func GetUserInfo(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	fmt.Println(vars)
	fmt.Fprint(res, vars["id"])
}

func CreatUserHandler(res http.ResponseWriter, req *http.Request) {
	// Create user structure first(pointer)
	user := new(User)

	// Use json parser & data flows to user *User
	err := json.NewDecoder(req.Body).Decode(user)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(res, err)
	}
	res.WriteHeader(http.StatusOK)
}

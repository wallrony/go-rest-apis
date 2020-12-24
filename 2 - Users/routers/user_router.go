package routers

import (
	"encoding/json"
	"strconv"

	"github.com/gorilla/mux"
	"gopkg.in/matryer/respond.v1"

	"../models"

	"net/http"
)

var users []models.User = []models.User{}

// IndexUsers Function returns all users saved in server instance
func IndexUsers(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode(users)
}

// ShowUser Function returns a specific User by id passed in request
// params.
func ShowUser(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	reqID, _ := strconv.ParseInt(params["id"], 10, 64)

	if reqID == 0 {
		respond.With(res, req, http.StatusBadRequest, map[string]string{
			"message": "need id param",
		})

		return
	}

	var user models.User

	for _, item := range users {
		if item.ID == reqID {
			user = item

			break
		}
	}

	if user.ID == 0 {
		respond.With(res, req, http.StatusNotFound, http.NoBody)
	} else {
		json.NewEncoder(res).Encode(user)
	}
}

// AddUser Function adds an User with body sended in request.
func AddUser(res http.ResponseWriter, req *http.Request) {
	var user models.User

	_ = json.NewDecoder(req.Body).Decode(&user)

	if user.Name == "" || user.Age == 0 {
		respond.With(res, req, http.StatusBadRequest, map[string]string{
			"message": "fields name or age cannot be empty",
		})

		return
	}

	user.ID = int64(len(users) + 1)

	users = append(users, user)

	respond.With(res, req, http.StatusCreated, users)
}

// DeleteUser Function deletes an user that ID is equal that the
// id passed in req params.
func DeleteUser(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	reqID, _ := strconv.ParseInt(params["id"], 10, 64)

	for index, user := range users {
		if user.ID == reqID {
			users = append(users[:index], users[index+1:]...)

			break
		}
	}

	return
}

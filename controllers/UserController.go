package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ASanchezT85/api-rest-mysql-go/commons"
	"github.com/ASanchezT85/api-rest-mysql-go/models"
	"github.com/gorilla/mux"
)

func Index(writer http.ResponseWriter, request *http.Request) {
	users := []models.User{}

	db := commons.GetConnection()
	defer db.Close()

	db.Find(&users)

	json, _ := json.Marshal(users)

	commons.SendResponse(writer, http.StatusOK, json)
}

func Show(writer http.ResponseWriter, request *http.Request) {
	user := models.User{}

	id := mux.Vars(request)["id"]

	db := commons.GetConnection()
	defer db.Close()

	db.First(&user, id)

	if user.ID > 0 {
		json, _ := json.Marshal(user)
		commons.SendResponse(writer, http.StatusOK, json)
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}

func Store(writer http.ResponseWriter, request *http.Request) {
	user := models.User{}
	db := commons.GetConnection()
	defer db.Close()

	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		commons.SendError(writer, http.StatusBadRequest)
		return
	}

	err = db.Save(&user).Error
	if err != nil {
		commons.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(user)
	commons.SendResponse(writer, http.StatusOK, json)
}

func Destroy(writer http.ResponseWriter, request *http.Request) {
	user := models.User{}

	db := commons.GetConnection()
	defer db.Close()

	id := mux.Vars(request)["id"]
	db.First(&user, id)
	if user.ID > 0 {
		db.Delete(&user)
		commons.SendResponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}

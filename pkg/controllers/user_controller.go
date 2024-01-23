package controllers

import (
	"log"
	"strconv"

	"encoding/json"
	"name_service/iternal/api"
	"name_service/iternal/models"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var Router *chi.Mux

/*
	post: http://localhost:8080

	request: {
		"name": "<NAME>", (required)
		"surname": "<Surname>", (required)
		"Patronymic": "<Patronymic>" (optional)
	}

	response: {
		"id": <id>
		"name": "<NAME>",
		"surname": "<Surname>",
		"Patronymic": "<Patronymic>",
		"age": "<age>",
		"gender": "<gender>",
		"nationality": "<nationality>"
	}
*/

func SaveNewName(responseWriter http.ResponseWriter, request *http.Request) {
	log.Printf("Get request for saving new name")
	var UserData models.UserData
	getDataFromRequest(responseWriter, request, &UserData)

	validateErr := ValidateSaveUserDataModel(&UserData)
	if JSONResponseException(validateErr, responseWriter) {
		log.Printf("Validation error: %s", validateErr.Error())
		return
	}

	savedUserData, serviceError := api.Save(&UserData)
	if JSONResponseException(serviceError, responseWriter) {
		log.Printf("Service error: %s", serviceError.Error())
		return
	}
	log.Printf("Save request for new name")
	JSONResponse(responseWriter, savedUserData)
}

/*
	put: http://localhost:8080

	request: {
		"id": <id> (required)
		"name": "<NAME>", (required)
		"surname": "<Surname>", (required)
		"Patronymic": "<Patronymic>" (optional)
		"age": <age> (optional)
		"gender": "<gender>" (optional)
		"nationality": "<nationality>" (optional)
	}

	response: {
		"id": <id>
		"name": "<NAME>",
		"surname": "<Surname>",
		"Patronymic": "<Patronymic>",
		"age": "<age>",
		"gender": "<gender>",
		"nationality": "<nationality>"
	}
*/

func UpdateUserData(responseWriter http.ResponseWriter, request *http.Request) {
	log.Printf("Get request for updating name")
	var UserData models.UserData
	getDataFromRequest(responseWriter, request, &UserData)

	validateErr := ValidateUpdateUserDataModel(&UserData)
	if JSONResponseException(validateErr, responseWriter) {
		log.Printf("Validation error: %s", validateErr.Error())
		return
	}

	updatedUserData, serviceError := api.Update(&UserData)
	if JSONResponseException(serviceError, responseWriter) {
		log.Printf("Service error: %s", serviceError.Error())
		return
	}
	log.Printf("Update request for name")
	JSONResponse(responseWriter, updatedUserData)
}

/*
	get: http://localhost:8080

	request: {
    "page": <PAGE>, (required)
    "user_filter": {
        "name": "<NAME>", (optional)
        "surname": "<SURNAME>", (optional)
        "patronymic": "<PATRONYMIC>", (optional)
        "gender": "<GENDER>", (optional)
        "nationality": "<NATIONALITY>" (optional)
    }

	response: {
		list:
			"id": <id>
			"name": "<NAME>",
			"surname": "<Surname>",
			"Patronymic": "<Patronymic>",
			"age": "<age>",
			"gender": "<gender>",
			"nationality": "<nationality>"
	}
*/

func GetUserData(responseWriter http.ResponseWriter, request *http.Request) {
	log.Printf("Get request for getting user data")

	var requestData models.UserDataRequest
	getDataFromRequest(responseWriter, request, &requestData)

	validateErr := ValidateGetUserDataModel(&requestData)
	if JSONResponseException(validateErr, responseWriter) {
		log.Printf("Validation error: %s", validateErr.Error())
		return
	}

	filteredUsers, serviceError := api.GetFilteredUsers(&requestData)
	if JSONResponseException(serviceError, responseWriter) {
		log.Printf("Service error: %s", serviceError.Error())
		return
	}

	log.Printf("Getting request was successful")

	JSONResponse(responseWriter, filteredUsers)
}

/*
	delete: http://localhost:8080/{userId}

	request: empty

	response: "User was deleted"
*/

func DeleteUserData(responseWriter http.ResponseWriter, request *http.Request) {
	log.Printf("Get request for deleting name by id")
	userID := chi.URLParam(request, "userID")

	userIDInt, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		JSONResponseException(err, responseWriter)
		return
	}

	serviceError := api.Delete(userIDInt)
	if JSONResponseException(serviceError, responseWriter) {
		log.Printf("Service error: %s", serviceError.Error())
		return
	}
	log.Printf("Delete request for deleting name by id")
	JSONResponse(responseWriter, "User was deleted")
}

func JSONResponse(responseWriter http.ResponseWriter, data interface{}) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	err := json.NewEncoder(responseWriter).Encode(data)
	JSONResponseException(err, responseWriter)
}

func getDataFromRequest[T any](responseWriter http.ResponseWriter, request *http.Request, object *T) {
	errorMapping := json.NewDecoder(request.Body).Decode(&object)
	JSONResponseException(errorMapping, responseWriter)
}

func JSONResponseException(err error, responseWriter http.ResponseWriter) bool {
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusBadRequest)
		return true
	}
	return false
}

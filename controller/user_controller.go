package controller

import (
	"encoding/json"
	"log"
	"net/http"
	servies "src/github.com/ns-saini/mvc/services"
	"src/github.com/ns-saini/mvc/utils"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userIdParam := req.URL.Query().Get("user_id")
	log.Printf("About to process for user %v ", userIdParam)
	user_id, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		json_value, _ := json.Marshal(&utils.ApplicationError{
			Message:    "user id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "Bad Request",
		})
		resp.Write(json_value)
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	user, userErr := servies.GetUser(user_id)
	if userErr != nil {
		json_value, _ := json.Marshal(userErr)
		resp.Write(json_value)
		resp.WriteHeader(userErr.StatusCode)
		return
	}

	json_value, _ := json.Marshal(user)
	resp.Write(json_value)

}

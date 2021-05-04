package domains

import (
	"log"
	"net/http"
	"src/github.com/ns-saini/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 1, FirstName: "Nishant", LastName: "Saini", Email: "nishtn.saini"},
	}
)

func GetUser(user_id int64) (*User, *utils.ApplicationError) {
	user := users[user_id]
	log.Printf("User: %v ", user)

	if user == nil {
		return nil, &utils.ApplicationError{
			Message:    "User not found",
			StatusCode: http.StatusNotFound,
			Code:       "Not found"}
	}
	return user, nil

}

package servies

import (
	"src/github.com/ns-saini/mvc/domains"
	"src/github.com/ns-saini/mvc/utils"
)

func GetUser(user_id int64) (*domains.User, *utils.ApplicationError) {
	return domains.GetUser(user_id)
}

package app

import (
	"net/http"
	"src/github.com/ns-saini/mvc/controller"
)

func StartApp() {

	http.HandleFunc("/users", controller.GetUser)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}

package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	io.WriteString(w, "Create User Handler")
// }

func RegisterHandels() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)

	router.POST("/user/:user_name", Login)

	return router
}

func main() {
	r := RegisterHandels()
	http.ListenAndServe(":8000", r)
}

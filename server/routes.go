package server

import (
	"github.com/julienschmidt/httprouter"
)

func setRoutes(route *httprouter.Router) {
	route.HandlerFunc("GET", "/user", logging(listUsers))
	route.HandlerFunc("POST", "/user", logging(createUser))
	route.HandlerFunc("DELETE", "/user", logging(deleteAllUsers))
	route.HandlerFunc("GET", "/user/:ID", logging(findUserByID))
	route.HandlerFunc("DELETE", "/user/:ID", logging(findAndDeleteUserByID))
}

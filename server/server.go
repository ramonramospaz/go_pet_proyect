package server

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//StartServer ...
func StartServer() {
	log.Println("Starting server")
	mux := httprouter.New()
	setRoutes(mux)
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	log.Printf("Addr:%v\n", server.Addr)
	server.ListenAndServe()
}

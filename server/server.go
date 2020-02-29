package server

import (
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

//StartServer ...
func StartServer() {
	log.Println("Starting server")
	mux := httprouter.New()
	setRoutes(mux)
	server := http.Server{
		Addr:    getAdressAndPort(),
		Handler: mux,
	}
	log.Printf("Addr:%v\n", server.Addr)
	server.ListenAndServe()
}

func getAdressAndPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8081"
		log.Println("INFO: No PORT enviroment variable, use default 8081 ")
	}
	return ":" + port
}

package server

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"twelve_webapp/controler"
)

func welcomePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("resources/templates/main.html"))
	tmpl.Execute(w, nil)
}

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("route:%v - %v", r.Method, r.URL.Path)
		f(w, r)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user controler.User
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	err := json.Unmarshal(body, &user)
	response := controler.Response{}
	// Always put the header before the WriteHeader
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	if err == nil {

		log.Println("Creating user")
		err = controler.CreateUser(user)
		if err == nil {
			response.Code = 200
			response.Message = "User Created"
		} else {
			log.Println("The user cant be created")
			response.Code = 400
			response.Message = "The fields; id, name, age, birthCountry are obligatory"
		}

	} else {

		log.Println("The user cant be created")
		response.Code = 400
		response.Message = "User cant be Created"
	}
	json.NewEncoder(w).Encode(response)
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	users := controler.ListUsers()
	// Always put the header before the WriteHeader
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(users)
}

func getParamFromURL(url string) (IDstr string) {
	routes := strings.Split(url, "/")
	IDstr = routes[len(routes)-1]
	return
}

func findUserByID(w http.ResponseWriter, r *http.Request) {
	IDstr := getParamFromURL(r.URL.String())
	ID, _ := strconv.Atoi(IDstr)
	user, err := controler.FindUserByID(ID)
	w.Header().Add("Content-Type", "application/json")
	if err == nil {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(user)
	} else {
		w.WriteHeader(400)
		response := controler.Response{}
		response.Code = 400
		response.Message = "User dont exist"
		json.NewEncoder(w).Encode(response)
	}
}

func findAndDeleteUserByID(w http.ResponseWriter, r *http.Request) {
	IDstr := getParamFromURL(r.URL.String())
	ID, _ := strconv.Atoi(IDstr)
	err := controler.FindAndDeleteUserByID(ID)
	response := controler.Response{}
	response.Code = 200
	response.Message = "The user deleted"
	if err != nil {
		response.Code = 400
		response.Message = "The user cant be deleted"
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)
}

func deleteAllUsers(w http.ResponseWriter, r *http.Request) {
	controler.DeleteAllUsers()
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	response := controler.Response{Code: 200, Message: "All uses where deleted"}
	json.NewEncoder(w).Encode(response)
}

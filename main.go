package main

import (
	"fmt"
	"hackathon/controllers"
	"hackathon/models"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "hackathon_dev"
)

const serverPort = ":3000"

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	svc, err := models.NewServices(psqlInfo)
	if err != nil {
		panic(err)
	}

	staticC := controllers.NewStatic()
	userC := controllers.NewUser(svc.User)

	r := mux.NewRouter()
	r.HandleFunc("/", staticC.Home)
	r.HandleFunc("/signup", staticC.SignUp).Methods("GET")
	r.HandleFunc("/login", staticC.Login).Methods("GET")
	r.HandleFunc("/signup", userC.Register).Methods("POST")

	fmt.Printf("Listening at port %s", serverPort)
	http.ListenAndServe(serverPort, r)
}
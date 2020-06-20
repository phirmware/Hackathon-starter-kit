package main

import (
	"fmt"
	"hackathon/controllers"
	"hackathon/middleware"
	"hackathon/models"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "hackathon_dev"
)

var serverPort = os.Getenv("PORT")

func main() {
	if serverPort == "" {
		serverPort = "3000"
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	svc, err := models.NewServices(psqlInfo)
	requireUserMW := middleware.NewRequireUser(svc.User)
	userMW := middleware.NewUser(svc.User)
	if err != nil {
		panic(err)
	}
	defer svc.Close()
	svc.AutoMigrate()
	// svc.DestroyAndCreate()

	staticC := controllers.NewStatic()
	userC := controllers.NewUser(svc.User)
	postC := controllers.NewPost(svc.Post)

	r := mux.NewRouter()
	r.HandleFunc("/", staticC.Home)
	r.HandleFunc("/signup", userC.SignUp).Methods("GET")
	r.HandleFunc("/signup", userC.Register).Methods("POST")
	r.HandleFunc("/login", userC.Login).Methods("GET")
	r.HandleFunc("/login", userC.SignIn).Methods("POST")
	r.HandleFunc("/cookie", userC.CookieTest).Methods("GET")
	r.HandleFunc("/protected", requireUserMW.RequireUserMiddleWare(userC.Protected)).Methods("GET")
	r.HandleFunc("/post", requireUserMW.RequireUserMiddleWare(postC.PostPage)).Methods("GET")
	r.HandleFunc("/post", requireUserMW.RequireUserMiddleWare(postC.HandlePost)).Methods("POST")
	r.HandleFunc("/list", requireUserMW.RequireUserMiddleWare(postC.ListPage)).Methods("GET")

	fmt.Printf("Listening at port %s", serverPort)
	http.ListenAndServe(":"+serverPort, userMW.UserMiddleWareFn(r))
}

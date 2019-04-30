package main

import (
	//"fmt"
	//"html/template"
	"Co-Lab/go_dev"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	//"time"

	"github.com/gorilla/mux"
)

//ROUTER CONTSRUCTER
//VERY IMPORTANT

var db *sql.DB

func newRouter() *mux.Router {
	r := mux.NewRouter()
	//http.HandleFunc("/favicon.ico", faviconHandler)
	//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	//THIS IS 100% VOODOO - DONT FUCKING TOUCH THIS UNDER ANY CIRCUMSTANCES
	//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))

	//ALL THE REST OF THESE ARE FINE TO MESS WITH
	//VIEWS SUB ROUTER
	s := r.PathPrefix("/view").Subrouter()
	s.HandleFunc("/", ViewHandler)
	s.HandleFunc("/{page}", ViewHandler)
	

	r.HandleFunc("/project_view/{key}", ProjectViewHandler)
	r.HandleFunc("/task_view/{key}", TaskViewHandler)

	//SESSIONS AND STUFF
	r.HandleFunc("/login", login)
	r.HandleFunc("/logout", logout)
	r.HandleFunc("/signup", signup)

	//DEFAULT ROUTE WHEN SOMEONE HITS THE SITE
	r.HandleFunc("/", IndexHandler)

	//404 HANDLEING WITH CUSTOM PAGE
	r.NotFoundHandler = http.HandlerFunc(notFound)

	return r
}

func main() {
	if debug == true {
		fmt.Println("Co-Lab core starting up")
	}

	db = go_dev.Initialize()
	if db == nil {
		fmt.Println("db is bad")
	}

	//WE NEED A ROUTER
	r := newRouter()

	port := ":8080"

	fmt.Println("go to ->  http://localhost" + port)
	//RUNS THE SERVER
	log.Fatal(http.ListenAndServe(port, r))
}

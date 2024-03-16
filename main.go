package main

import (
	"link-walker-go/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/task", controller.GetTasks).Methods("GET")
	router.HandleFunc("/task", controller.PostTask).Methods("POST")

	println("Started server at port: 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		println("Failed to start server")
	}
}



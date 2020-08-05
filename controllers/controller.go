package controllers

import (
	"log"
	"net/http"
	//"github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("conroller..")
	w.Write([]byte("Gorilla!\n"))
}

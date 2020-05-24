package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my about page! </h1>")
}

func StartServer() {
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	http.ListenAndServe(":3000", router)
}

func main() {
	StartServer()
}

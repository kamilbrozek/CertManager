package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func get(responseWritter http.ResponseWriter, request *http.Request) {
	responseWritter.Header().Set("Content-Type", "text/html")
	responseWritter.WriteHeader(http.StatusOK)
	responseWritter.Write([]byte(`Get Called`))
}

func main() {
	fmt.Printf("Test")
	muxRouter := mux.NewRouter()
	api := muxRouter.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", get).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", muxRouter))

}

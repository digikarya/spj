package main

import (
	"github.com/digikarya/spj/app"
	"github.com/digikarya/spj/config"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)
// ServeHTTP wraps the HTTP server enabling CORS headers.
// For more info about CORS, visit https://www.w3.org/TR/cors/
func main() {
	r := mux.NewRouter()
	conf := config.GetConfig()
	SPJ := &app.SPJ{}
	SPJ.Initialize(conf,r)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Print("App is not running")
	}else{
		log.Print("App is not running")
	}
}


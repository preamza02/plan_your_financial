package main

import (
	"fmt"
	"log"
	"net/http"
	"plan_your_financial/internal/backend"
	"plan_your_financial/internal/frontend"
)

func main(){
	port:= "8080"
	mux := http.NewServeMux()
	frontend.InitFrontend(mux)
	backend.InitBackend(mux)
	log.Printf("Server starting on :%s\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s",port), nil)
	if err != nil {
		log.Fatal(err)
	}
}


package hello 

import (
	"encoding/json"
	"net/http"
)

type HelloResponse struct {
	Message string `json:"message"`
}

func getHelloService() HelloResponse {
	return HelloResponse{Message: "Hello, World! This is Working"}
}

func GetHelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	resp := getHelloService()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

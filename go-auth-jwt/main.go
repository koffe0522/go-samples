package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	auth "github.com/go-auth-jwt/src"
	"github.com/joho/godotenv"
)

type response struct {
	Result string `json:"result"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/login.html")
	if err != nil {
		log.Fatalf("Not found file: %v", err)
	}

	t.Execute(w, nil)
}

var private = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	response := &response{
		Result: "success",
	}
	json.NewEncoder(w).Encode(response)
})

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	loadEnv()
	http.HandleFunc("/auth", auth.GetTokenHandler)
	http.Handle("/private", auth.JwtMiddleware.Handler(private))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("error starting http server: ", err)
		return
	}
}

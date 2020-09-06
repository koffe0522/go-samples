package main

import (
	"log"
	"net/http"
	"text/template"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/hello.html")
	if err != nil {
		log.Fatalf("Not found file: %v", err)
	}

	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("error starting http server: ", err)
		return
	}
}

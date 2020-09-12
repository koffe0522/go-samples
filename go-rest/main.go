package main

import (
	"koffe0522/go-projects/go-rest/database"
	"koffe0522/go-projects/go-rest/routes"
	"log"
	"net/http"
)

func main() {
	database.Init()
	defer database.Db.Close()
	routes := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", routes))
}

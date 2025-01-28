package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/Utsav_Kasvala/Netflix_backend/routes"
)

func main() {

	r := router.Router()
	fmt.Println("Starting server on the port 4000...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Server started on the port 4000...")
}

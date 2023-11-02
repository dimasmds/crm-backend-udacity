package main

import (
	"fmt"
	"github.com/dimasmds/crm-backend-udacity/pkg/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.GetConsumersRouter()
	fmt.Println("server start at http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}

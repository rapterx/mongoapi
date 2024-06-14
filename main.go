package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rapterx/mongoapi/router"
)

func main() {

	r := router.Router()

	fmt.Println("Server is getting started...")

	log.Fatal(http.ListenAndServe(":4000", r))

	fmt.Println("Listening at port 4000 ...")
}

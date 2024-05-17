package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kiruiaaron/go-postgres-yt/router"
)



func main(){
	r := router.Router()

	fmt.Println("Starting the server on the port 8080 ...")

	log.Fatal(http.ListenAndServe(":8080", r))


}

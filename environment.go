package main

import (
	"fmt"
	"log"
	"net/http"
	"google.golang.org/appengine"
)


func Handler (w http.ResponseWriter, r *http.Request) {

	cloudengine :=  appengine.IsDevAppServer()
	appserver := appengine.IsAppEngine()

	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	fmt.Fprintf(w, "Dev %t, App %t \n", cloudengine, appserver)
}

func main() {

	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe( ":8081", nil ))

}

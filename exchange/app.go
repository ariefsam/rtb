package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ariefsam/godotenv"
	"github.com/ariefsam/pure"
	"github.com/ariefsam/rtb/exchange/router"
)

func main() {

	var config map[string]string
	config, err := godotenv.Read()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("APP Name:", config["APP_NAME"])

	p := pure.New()

	p.Get("/", home)
	router.Register(p)

	srv := &http.Server{
		Handler: p.Serve(),
		Addr:    "localhost:8011",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 150 * time.Second,
		ReadTimeout:  150 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome")
}

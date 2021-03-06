package main

import (
	"fmt"
	"go-heroku-example/pkg/transport"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	port := os.Getenv("PORT")

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Println("200", err)
	} else {
		fmt.Println("Fatal, error")
	}

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	r := transport.Router()
	fmt.Println(http.ListenAndServe(":"+port, r))

}

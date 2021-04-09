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

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	r := transport.Router()
	fmt.Println(http.ListenAndServe(":"+port, r))

}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	tittle string
	price  string
}

func main() {

	r := gin.Default()

	svr := http.Server{
		Addr:        "8080",
		Handler:     r,
		ReadTimeout: 10,
	}

}

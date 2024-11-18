package main

import (
	"net/http"

	handler "github.com/mateoops/linkoln/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	handler.NewHandler(&handler.Config{
		R: router,
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	srv.ListenAndServe()
}

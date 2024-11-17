package main

import (
	handler "mateoops/linkoln/handlers"
	"net/http"

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

package handler

import (
	short "mateoops/linkoln/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

type Config struct {
	R *gin.Engine
}

func NewHandler(c *Config) {
	h := &Handler{}

	g := c.R.Group("/short")

	g.POST("/", h.Encode)
}

func (h *Handler) Encode(c *gin.Context) {
	shortUrl := short.Short{}
	if err := c.BindJSON(&shortUrl); err != nil {
		return
	}

	// bellow tests

	// set ID
	shortUrl.ID = 5
	// set short
	shortUrl.ShortUrl = "aksfjghfk"
	// set created
	shortUrl.CreatedAt = "today"
	// set Views to 0
	shortUrl.Views = 0
	c.IndentedJSON(http.StatusCreated, shortUrl)
}

// func Decode

// func Redirect

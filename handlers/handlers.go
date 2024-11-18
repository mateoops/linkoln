package handlers

import (
	"context"
	"mateoops/linkoln/models"
	"mateoops/linkoln/repositories"
	"mateoops/linkoln/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *services.ShortService
}

type Config struct {
	R *gin.Engine
}

func NewHandler(c *Config) {
	h := &Handler{}

	g := c.R.Group("/short")

	g.POST("/", h.Encode)
	g.GET("/:shortPath", h.Decode)

	// Database DI
	mongoShortRepo := repositories.NewMongoShortRepo()
	h.svc = services.NewShortService(mongoShortRepo)
}

func (h *Handler) Encode(c *gin.Context) {
	short := models.Short{}
	if err := c.BindJSON(&short); err != nil {
		return
	}

	// bellow tests

	// set ID
	short.ID = 5
	// set short
	short.ShortUrl = "aksfjghfk"
	// set created
	short.CreatedAt = "today"
	// set Views to 0
	short.Views = 0

	shortURL, err := h.svc.CreateShort(context.TODO(), short)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, short)
	} else {
		c.IndentedJSON(http.StatusCreated, shortURL)
	}

}

// func Decode
func (h *Handler) Decode(c *gin.Context) {
	shortPath := c.Param("shortPath")
	short := h.svc.GetByShortUrl(context.TODO(), shortPath)
	if short.Url == "" {
		c.IndentedJSON(http.StatusNotFound, "Object not found")
	} else {
		c.IndentedJSON(http.StatusOK, short)
	}
}

// func Redirect

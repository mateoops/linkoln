package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mateoops/linkoln/internal"
	"github.com/mateoops/linkoln/models"
	"github.com/mateoops/linkoln/repositories"
	"github.com/mateoops/linkoln/services"

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

	// set ID
	short.ID = uuid.NewString()
	// set short
	short.ShortUrl = internal.GenerateShortID(6)
	// set created
	short.CreatedAt = time.Now().String()

	shortURL, err := h.svc.CreateShort(context.TODO(), short)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, short)
	} else {
		c.IndentedJSON(http.StatusCreated, gin.H{"shortUrl": shortURL})
	}

}

// func Decode
func (h *Handler) Decode(c *gin.Context) {
	shortPath := c.Param("shortPath")
	short := h.svc.GetByShortUrl(context.TODO(), shortPath)
	if short.Url == "" {
		c.IndentedJSON(http.StatusNotFound, "Object not found")
	} else {
		//c.IndentedJSON(http.StatusOK, short)
		c.Redirect(http.StatusFound, short.ShortUrl)
	}
}

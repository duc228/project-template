package controllers

import (
	"net/http"
	"server/internal/db"
	m "server/internal/db/models"
	"server/pkg/app"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	app.NewResponse(c, http.StatusOK, db.Albums)
}

func PostAlbums(c *gin.Context) {
	var newAlbum m.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	db.Albums = append(db.Albums, newAlbum)
	app.NewResponse(c, http.StatusCreated, db.Albums)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range db.Albums {
		if a.ID == id {
			app.NewResponse(c, http.StatusOK, a)
			return
		}
	}
	app.NewResponse(c, http.StatusNotFound, nil)
}

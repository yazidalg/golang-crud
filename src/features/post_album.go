package features

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	s "golang-crud/src/structs"
)

func AddAlbum(c *gin.Context) {
	var newAlbum s.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		log.Fatal(err)
		return
	}

	s.Albums = append(s.Albums, newAlbum)
	c.JSON(http.StatusCreated, gin.H{
		"message":    "success",
		"statusCode": http.StatusCreated,
		"data":       newAlbum,
	})
}

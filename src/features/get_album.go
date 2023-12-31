package features

import (
	"net/http"

	"github.com/gin-gonic/gin"

	s "golang-crud/src/structs"
)

func GetAlbums(c *gin.Context) {
	var albums = s.Albums

	c.JSON(http.StatusOK, gin.H{
		"message":    "success",
		"statusCode": http.StatusOK,
		"data":       albums,
	})
}

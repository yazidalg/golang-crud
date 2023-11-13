package features

import (
	"net/http"

	"github.com/gin-gonic/gin"

	s "golang-crud/src/structs"
)

func GetAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range s.Albums {
		if a.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"message":    "success",
				"statusCode": http.StatusOK,
				"data":       a,
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message":    "Not Found",
		"statusCode": http.StatusNotFound,
		"data":       nil,
	})
}

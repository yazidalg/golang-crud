package main

import (
	"fmt"

	f "golang-crud/src/features"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("CRUD With Golang")

	router := gin.Default()
	router.GET("/albums", f.GetAlbums)
	router.POST("/albums", f.AddAlbum)

	router.Run("localhost:8080")
}

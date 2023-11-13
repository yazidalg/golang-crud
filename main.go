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

	router.Run("localhost:8080")
}

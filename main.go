package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "👋 Hi! I'm Ginapp")
	})

	r.GET("welcome", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": "true"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatal("cannot initialize server")
	}
}

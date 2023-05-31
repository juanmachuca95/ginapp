package main

import (
	sqlDB "ginapp/db"
	"ginapp/handlers"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := sqlDB.ConexionSql()
	if err != nil {
		panic(err)
	}

	// questions
	questions := handlers.NewQuestions(db)
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "👋 Hi! I'm Ginapp")
	})

	v1 := r.Group("v1")
	v1.GET("/questions", questions.Get())
	v1.POST("/questions", questions.Save())
	v1.PUT("/questions", questions.Update())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatal("cannot initialize server")
	}
}

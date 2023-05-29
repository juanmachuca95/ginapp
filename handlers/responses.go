package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseRequest struct {
	Message string `json:"message"`
}

type Responses struct {
	db *sql.DB
}

func NewResponses(db *sql.DB) Responses {
	return Responses{
		db: db,
	}
}

func (r Responses) Save() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ResponseRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		// TODO: Save response and
		log.Println("Message ", req.Message)
		c.Status(http.StatusOK)
	}
}

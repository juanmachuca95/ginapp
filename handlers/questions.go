package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QuestionRequest struct {
	App string `form:"app"`
}

type QuestionResponse struct {
	Question string         `json:"question"`
	Status   QuestionStatus `json:"status"` // answered or unanswered
}

type QuestionStatus string

const (
	answered   QuestionStatus = "answered"
	unanswered QuestionStatus = "unanswered"
)

type Questions struct {
	db *sql.DB
}

func NewQuestions(db *sql.DB) Questions {
	return Questions{
		db: db,
	}
}

func (q Questions) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req QuestionRequest
		if err := c.ShouldBindQuery(&req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		// TODO: Get last question

		log.Println(req.App)
		c.JSON(http.StatusOK, QuestionResponse{
			Question: "Puedes comentarme que enfoque le dio la serie Crown a lo sucedido con las islas malvinas?",
			Status:   unanswered,
		})
	}
}

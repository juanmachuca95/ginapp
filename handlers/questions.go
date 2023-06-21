package handlers

import (
	"database/sql"
	"ginapp/db"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type QuestionRequest struct {
	App string `form:"app"`
}

type ResponseRequest struct {
	QuestionID uint   `json:"question_id"`
	Answer     string `json:"answer"`
}

type SaveRequest struct {
	Question string         `json:"question"`
	Status   QuestionStatus `json:"status"`
}

type QuestionResponse struct {
	ID       uint           `json:"id"`
	Question string         `json:"question"`
	Status   QuestionStatus `json:"status"` // answered or unanswered
}

type AnswerResponse struct {
	Answer string         `json:"answer"`
	Status QuestionStatus `json:"status"`
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
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		stmt, err := q.db.PrepareContext(c, db.QUESTION_LAST_UNANSWERED)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		defer stmt.Close()

		var resp QuestionResponse
		if err := stmt.QueryRowContext(c).Scan(&resp.ID, &resp.Question, &resp.Status); err != nil {
			if err == sql.ErrNoRows {
				c.Status(http.StatusOK)
				return
			}
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, resp)
	}
}

func (q Questions) Save() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req SaveRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		stmt, err := q.db.PrepareContext(ctx, db.QUESTION_SAVE)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		defer stmt.Close()

		result, err := stmt.ExecContext(ctx, req.Question, unanswered)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		rowsAffected, _ := result.RowsAffected()
		rowID, _ := result.LastInsertId()
		log.Println("Rows affected: ", rowsAffected, rowID)

		duration := 120 * time.Second
		timeout := time.After(duration)

		var response AnswerResponse
		for {
			select {
			case <-timeout:
				ctx.AbortWithStatus(http.StatusRequestTimeout)
				return

			default:
				time.Sleep(time.Second * 5)
				stmt, err := q.db.PrepareContext(ctx, db.QUESTION_BY_ID)
				if err != nil {
					ctx.AbortWithError(http.StatusInternalServerError, err)
					return
				}
				defer stmt.Close()

				if err := stmt.QueryRowContext(ctx, rowID).Scan(&response.Answer, &response.Status); err != nil {
					if err == sql.ErrNoRows {
						continue
					}
					ctx.AbortWithError(http.StatusInternalServerError, err)
					return
				}

				ctx.JSON(http.StatusOK, response)
				return
			}
		}
	}
}

func (q Questions) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ResponseRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		stmt, err := q.db.PrepareContext(c, db.UPDATE_QUESTION)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		defer stmt.Close()

		_, err = stmt.ExecContext(c, req.Answer, answered, req.QuestionID)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Status(http.StatusOK)
	}
}

package main

import (
	"encoding/json"
	"errors"
	"github.com/KevinZonda/float32/exec/svr/db"
	"github.com/KevinZonda/float32/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ginHistory(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		utils.GinErrorMsg(c, errors.New("id not set"))
		return
	}

	ans, err := db.FindAnswerById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{
				"message": "not found",
			})
			return
		}
		utils.GinErrorMsg(c, err)
		return
	}

	var evi []SearchItem
	if ans.Evidence != "" {
		_ = json.Unmarshal([]byte(ans.Evidence), &evi)
	}
	c.JSON(200, HistoryResponse{
		Question: ans.Question,
		Answer:   ans.FirstAnswer,
		Evidence: evi,
	})

}

type HistoryResponse struct {
	Question string       `json:"question"`
	Answer   string       `json:"answer"`
	Evidence []SearchItem `json:"evidence"`
}

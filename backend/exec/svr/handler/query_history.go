package handler

import (
	"encoding/json"
	"errors"
	"github.com/KevinZonda/float32/exec/svr/db"
	"github.com/KevinZonda/float32/exec/svr/reqmodel"
	"github.com/KevinZonda/float32/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func QueryHistory(c *gin.Context) {
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
		utils.GinErrorMsg(c, errors.New("db error"))
		return
	}

	var evi []reqmodel.SearchItem
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
	Question string                `json:"question"`
	Answer   string                `json:"answer"`
	Evidence []reqmodel.SearchItem `json:"evidence"`
}

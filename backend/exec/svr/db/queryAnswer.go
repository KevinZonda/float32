package db

import (
	"errors"
	"github.com/KevinZonda/float32/exec/svr/dbmodel"
	"math/rand"
)

func FindAnswerById(id string) (answer dbmodel.Answer, err error) {
	if _db == nil {
		return dbmodel.Answer{}, errors.New("history not available")
	}
	err = _db.First(&answer, "id = ?", id).Error
	return
}

func NewAnswer(ans dbmodel.Answer) (dbmodel.Answer, error) {
	if _db == nil {
		return ans, errors.New("history not available")
	}
	if ans.ID == "" {
		ans.ID = rndStr(10)
	}

	err := _db.Create(&ans).Error
	return ans, err
}

func UpdateAnswer(ans dbmodel.Answer) (dbmodel.Answer, error) {
	if _db == nil {
		return dbmodel.Answer{}, errors.New("history not available")
	}
	err := _db.Save(&ans).Error
	return ans, err
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func rndStr(lens int) string {
	b := make([]byte, lens)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

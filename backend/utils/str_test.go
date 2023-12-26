package utils_test

import (
	"github.com/KevinZonda/float32/utils"
	"testing"
)

func TestStrC(t *testing.T) {

	ss := []string{"afheuhfweyuwefuyh", "bwecfuih中文", "欧欧欧欧欧欧欧"}
	for _, s := range ss {
		t.Log(s, utils.StrByteRuneDiffRate(s))

	}
}

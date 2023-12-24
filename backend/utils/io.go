package utils

import (
	"encoding/json"
	"io"
)

func ReadToEndJson[T any](r io.Reader) (T, error) {
	var t T
	if r == nil {
		return t, io.EOF
	}
	err := json.NewDecoder(r).Decode(&t)
	return t, err
}

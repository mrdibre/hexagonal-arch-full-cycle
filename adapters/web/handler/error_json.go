package handler

import (
	"encoding/json"
	"fmt"
)

func jsonError(msg string) []byte {
	e := struct {
		Message string `json:"message"`
	}{msg}

	r, err := json.Marshal(e)

	if err != nil {
		return []byte(err.Error())
	}

	fmt.Println(e)
	return r
}

package handler

import "encoding/json"

func jsonError(msg string) []byte {
	e := struct {
		message string `json:"message"`
	}{ msg }

	r, err := json.Marshal(e)

	if err != nil {
		return []byte(err.Error())
	}

	return r
}

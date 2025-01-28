package handler

import "encoding/json"

func jsonError(msg string) []byte {
	errorMsg := struct {
		Message string `json:"message"`
	}{
		msg,
	}

	r, err := json.Marshal(errorMsg)
	if err != nil {
		return []byte(err.Error())
	}
	return r
}

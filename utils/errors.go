package utils

import "net/http"

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NotFound() Error {
	e := Error{Status: http.StatusNotFound, Message: "Not found"}
	return e
}

package common

import "net/http"

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

var (
	NotFound = Error{Status: http.StatusNotFound, Message: "Not found"}
)

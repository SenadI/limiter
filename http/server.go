package http

import (
	"net/http"
	"time"
)

// NewGreetingsHandler - Returns configured router object with greetings handler attached.
func NewGreetingsHandler() http.HandlerFunc {
	return greetingsHandler(time.Now())
}

package main

import (
	"fmt"
	"net/http"
)

// Handler responsible for returning 1x1 gif
func NewImgHandler() http.Handler {
	return &imgHandler{}
}

type imgHandler struct{}

const gif1x1 = "\x47\x49\x46\x38\x39\x61\x01\x00\x01\x00\x80\x00\x00\x00\x00" +
	"\x00\xFF\xFF\xFF\x21\xF9\x04\x01\x00\x00\x00\x00\x2C\x00\x00\x00\x00" +
	"\x01\x00\x01\x00\x00\x02\x01\x44\x00\x3B"

func (*imgHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/gif")
	fmt.Fprintf(w, gif1x1)
}

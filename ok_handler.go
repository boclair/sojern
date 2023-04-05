package main

import (
	"fmt"
	"net/http"
)

// Handler responsible for returning OK, or 503 based on boolean status
func NewOkHandler() OkHandler {
	return &okHandlerImpl{}
}

type OkHandler interface {
	http.Handler
	setStatus()
	clearStatus()
}

type okHandlerImpl struct {
	status bool
}

func (h *okHandlerImpl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.status {
		fmt.Fprintf(w, "OK")
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
}

func (h *okHandlerImpl) setStatus() {
	h.status = true
}

func (h *okHandlerImpl) clearStatus() {
	h.status = false
}

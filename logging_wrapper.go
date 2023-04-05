package main

import (
	"log"
	"net/http"
)

const logBufferSize = 1000

// Factory used to wrap an http.Handler to add logging.
// Factory automatically creates the go routine to read from logging channel.
func NewLoggingWrapper() LoggingWrapper {
	loggingChan := make(chan logFormat, logBufferSize)
	go logReader(loggingChan)

	return &wrapperImpl{loggingChan: loggingChan}
}

// Presents the wrap method to wrap any http.Handler with logging functionality.
type LoggingWrapper interface {
	wrap(nextHandler http.Handler) http.Handler
}

func logReader(loggingChan <-chan logFormat) {
	for entry := range loggingChan {
		log.Println(entry)
	}
}

// Logged structure on every log wrapped request.
type logFormat struct {
	host       string
	path       string
	referer    string
	remoteAddr string
	userAgent  string
}

type wrapperImpl struct {
	loggingChan chan<- logFormat
}

func (h *wrapperImpl) wrap(nextHandler http.Handler) http.Handler {
	return &wrappedHandler{
		loggingChan: h.loggingChan,
		nextHandler: nextHandler,
	}
}

type wrappedHandler struct {
	loggingChan chan<- logFormat
	nextHandler http.Handler
}

func (h *wrappedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := logFormat{
		host:       r.Host,
		path:       r.URL.Path,
		referer:    r.Referer(),
		remoteAddr: r.RemoteAddr,
		userAgent:  r.UserAgent(),
	}

	select {
	case h.loggingChan <- msg:
		// Success
	default:
		// TODO: Log buffer is full. Just ignoring for now.
	}
	h.nextHandler.ServeHTTP(w, r)
}

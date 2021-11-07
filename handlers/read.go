package handlers

import (
	"log"
	"net/http"
)

type ReadImages struct {
	l *log.Logger
}

func (gh *ReadImages) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("GET"))
}

func NewReadImages(l *log.Logger) *ReadImages {
	return &ReadImages{l}
}

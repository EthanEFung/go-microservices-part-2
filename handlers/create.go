package handlers

import (
	"log"
	"net/http"
)

type CreateImage struct {
	l *log.Logger
}

func (ch *CreateImage) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("POST"))
}

func NewCreateImage(l *log.Logger) *CreateImage {
	return &CreateImage{l}
}

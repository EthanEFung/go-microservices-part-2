package handlers

import (
	"log"
	"net/http"
)

type UpdateImage struct {
	l *log.Logger
}

func (uh *UpdateImage) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("PUT"))
}

func NewUpdateImage(l *log.Logger) *UpdateImage {
	return &UpdateImage{l}
}

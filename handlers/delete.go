package handlers

import (
	"log"
	"net/http"
)

type DeleteImage struct {
	l *log.Logger
}

func (dh *DeleteImage) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("DELETE"))
}

func NewDeleteImage(l *log.Logger) *DeleteImage {
	return &DeleteImage{l}
}

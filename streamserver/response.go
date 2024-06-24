package main

import (
	"io"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, errMsg string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, errMsg)
}

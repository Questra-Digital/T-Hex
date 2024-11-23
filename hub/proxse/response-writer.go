package main

import (
	"bytes"
	"net/http"
)

// http.ResponseWriter but saves statusCode and body
type ResponseWriterSaver struct {
	http.ResponseWriter
	statusCode int
	body       bytes.Buffer
}

func (w *ResponseWriterSaver) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *ResponseWriterSaver) Write(data []byte) (int, error) {
	if w.statusCode == 0 {
		w.statusCode = http.StatusOK
	}
	w.body.Write(data)
	return w.ResponseWriter.Write(data)
}

// StatusCode returns the saved HTTP status code.
func (w *ResponseWriterSaver) StatusCode() int {
	return w.statusCode
}

// Body returns the captured response body as a string.
func (w *ResponseWriterSaver) Body() string {
	return w.body.String()
}

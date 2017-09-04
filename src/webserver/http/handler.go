package http

import (
	"fmt"
	"net/http"
)

type TextHandler struct {
	Response string
}

func (th *TextHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, th.Response)
}

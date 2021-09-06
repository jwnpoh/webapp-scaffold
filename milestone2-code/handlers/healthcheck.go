package handlers

import (
	"fmt"
	"net/http"
)

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
	return
}

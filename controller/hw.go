package controller

import (
	"fmt"
	"net/http"
)

func Helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "service is up!")
}

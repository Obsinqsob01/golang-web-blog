package actions

import (
	"fmt"
	"net/http"
)

// HomeHandler home route
func HomeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
}

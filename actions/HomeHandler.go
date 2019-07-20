package actions

import (
	"html/template"
	"net/http"
)

// HomeHandler home route
func HomeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))

		data := make(map[string]interface{})

		data["Title"] = "GoLang Blog"

		err := tmpl.Execute(w, data)
		if err != nil {
			panic(err)
		}
	})
}

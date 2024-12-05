package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/aaron-g-sanchez/PROJECTS/WORKOUT-TRACKER/views"
)

func main() {
	view := views.Index()

	http.Handle("/", templ.Handler(view))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}

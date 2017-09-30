package handlers

import "net/http"

// Index is the starting point for the web application from an end user
// perspective.
func Index(w http.ResponseWriter, r *http.Request) {
	render(w, "index")
}

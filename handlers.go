package main

import "net/http"

// check if the user is logged in and display posts accordingly
// handle post filtering by category and mode
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		showError(w, 400, "400 BAD REQUEST")
	}
	
}
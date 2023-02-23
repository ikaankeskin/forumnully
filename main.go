package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var (
	SESSION_ID = "SESSION_ID"
	port       = "8080"
)

/*    		End points
/    					main page handler index.html
/sign					signIn/signUp handler sign.html
/signup					register user handler
/login					log in handler
/signout				log out handler
/post					new post handler post.html
/savepost				save post handler
/registerlike			save post like handler
/registercommentlike	save comment like handler
/comment				new comment comment.html
/commentsubmit			save comment handler
/setfilter				add filter by category handler
/removefilter			remove filter by category handler
/changemode				change view mode handler. View modes: Shaw All/My Posts/My Comments/My Likes
*/

func main() {
	DB, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal(err)
	}
	db = DB

	defer db.Close()

	createTables()

	http.HandleFunc("/", homeHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	fmt.Println("Server starts at port :8080")
	http.ListenAndServe(":"+port, nil)
}

func createTables() {
	err := crerateUsersTable()
	if err != nil {
		log.Fatal(err)
	}
	err = crerateCategoriesTable()
	if err != nil {
		log.Fatal(err)
	}
	err = creratePostsTable()
	if err != nil {
		log.Fatal(err)
	}
	err = crerateCommentsTable()
	if err != nil {
		log.Fatal(err)
	}
	err = creratePostLikesTable()
	if err != nil {
		log.Fatal(err)
	}
	err = crerateCommentLikesTable()
	if err != nil {
		log.Fatal(err)
	}
}

func showError(w http.ResponseWriter, code int, message string) {
	// set the status code of the HTTP response to the provided 'code' int
	// we do it before excuting the template to ensure the correct HTTP status code is included in the response headers
	w.WriteHeader(code)
	// 	// execute the parsed templ by calling the Ecexute method,
	// passing the w and message args, substituting the {{.}} with the message arg

	templ, err := template.ParseFiles("templates/error.html")
	if err != nil {
		fmt.Fprint(w, "500 Internal Error")
		return
	}
	err = templ.Execute(w, message)
	if err != nil {
		fmt.Fprint(w, "500 Internal Error")
	}
}

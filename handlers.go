package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--Users---------")
	printUsers()
	fmt.Println("--Categories----")
	printCategories()
	fmt.Println("--Posts---------")
	printPosts()
	fmt.Println("--Comments------")
	printComments()
	fmt.Println("--Post Likes----")
	printPostLikes()
	fmt.Println("--Comment Likes-")
	printCommentLikes()
	invalidCredentialsFlagSignUp = false
	invalidCredentialsFlagSignIn = false
	emptyPostFlag = false
	indexObject := IndexObject{}
	sessionId := getCookie(r)
	indexObject.User = getUserBySessionId(sessionId)
	posts, err := getPosts(indexObject.User)
	if err != nil {
		showError(w, 500, "500 Internal Server Error")
		return
	}
	posts = filterByCategories(posts, filterCategories)
	indexObject.Posts = posts
	indexObject.Filters = filterCategories
	templ, err := template.ParseFiles("templates/index.html")
	if err != nil {
		showError(w, 500, "500 Internal Server Error")
		return
	}
	err = templ.Execute(w, indexObject)
	if err != nil {
		showError(w, 500, "500 Internal Server Error")
		return
	}
}
func signHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("templates/sign.html")
	if err != nil {
		showError(w, 500, "500 Internal Server Error")
		return
	}
	singCredentials := SingCredentials{}
	singCredentials.SignIn = invalidCredentialsFlagSignIn
	singCredentials.SignUp = invalidCredentialsFlagSignUp
	err = templ.Execute(w, singCredentials)
	if err != nil {
		showError(w, 500, "500 Internal Server Error")
		return
	}
}
func signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		showError(w, 405, "405 Method Not Allowed")
		return
	}
	username := r.FormValue("signup_username")
	email := r.FormValue("signup_email")
	password := r.FormValue("signup_password")
	sessionId := generateSessionId()
	err := saveUser(username, email, encrypt(password), sessionId)
	if err != nil {
		if strings.HasPrefix(err.Error(), "UNIQUE constraint failed:") {
			invalidCredentialsFlagSignUp = true
			invalidCredentialsFlagSignIn = false
			http.Redirect(w, r, "/sign", http.StatusTemporaryRedirect)
			return
		}
		showError(w, 500, "500 Internal Server Error. Error while working with database")
		return
	}
	setCookie(w, sessionId)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		showError(w, 405, "405 Method Not Allowed")
		return
	}
	email := r.FormValue("login_email")
	password := r.FormValue("login_password")
	user, err := checkUser(email, password)
	if err != nil {
		showError(w, 500, "500 Internal Server Error. Error while working with database")
		return
	}
	if err == nil && user == nil {
		invalidCredentialsFlagSignIn = true
		invalidCredentialsFlagSignUp = false
		http.Redirect(w, r, "/sign", http.StatusTemporaryRedirect)
		return
	}
	if user != nil {
		sessionId := generateSessionId()
		setCookie(w, sessionId)
		err := setSessionId(user, sessionId)
		if err != nil {
			showError(w, 500, "500 Internal Server Error")
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
func signoutHandler(w http.ResponseWriter, r *http.Request) {
	sessionId := getCookie(r)
	if sessionId != "" {
		err := resetSessionId(sessionId)
		if err != nil {
			showError(w, 500, "500 Internal Server Error. Error while working with database")
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
func postHandler(w http.ResponseWriter, r *http.Request) {
	sessionId := getCookie(r)
	user := getUserBySessionId(sessionId)
	if user == nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	templ, err := template.ParseFiles("templates/post.html")
	if err != nil {
		showError(w, 500, "500 Internal Server Error")
		return
	}
	newPostObject := NewPostObject{}
	newPostObject.User = user
	categories, err := getCategories()
	if err != nil {
		showError(w, 500, "500 Internal Server Error")
		return
	}
	newPostObject.Categories = categories
	newPostObject.IsEmptyPost = emptyPostFlag
	err = templ.Execute(w, newPostObject)
	if err != nil {
		showError(w, 500, "500 Internal Server Error")
		return
	}
}
func savepostHandler(w http.ResponseWriter, r *http.Request) {
	sessionId := getCookie(r)
	user := getUserBySessionId(sessionId)
	if user == nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	if r.Method != "POST" {
		showError(w, 405, "405 Method Not Allowed")
		return
	}
	postContent := strings.TrimSpace(r.FormValue("post_content"))
	postCategories := r.FormValue("categories")
	if postContent == "" {
		emptyPostFlag = true
		http.Redirect(w, r, "/post", http.StatusTemporaryRedirect)
		return
	}
	err := savePost(*user, postContent, postCategories)
	if err != nil {
		showError(w, 500, "500 Internal Server Error")
		return
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
func registerlikeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		sessionId := getCookie(r)
		user := getUserBySessionId(sessionId)
		if user == nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		postIdStr := r.FormValue("postId")
		statusStr := r.FormValue("status")
		postId, err := strconv.Atoi(postIdStr)
		if err != nil {
			return
		}
		status, err := strconv.Atoi(statusStr)
		if err != nil {
			return
		}
		updatePostLikes(user, postId, status)
	}
}
func registercommentlikeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		sessionId := getCookie(r)
		user := getUserBySessionId(sessionId)
		if user == nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		commentIdStr := r.FormValue("commentId")
		statusStr := r.FormValue("status")
		commentId, err := strconv.Atoi(commentIdStr)
		if err != nil {
			return
		}
		status, err := strconv.Atoi(statusStr)
		if err != nil {
			return
		}
		updatePostCommentLikes(user, commentId, status)
	}
}
func commentHandler(w http.ResponseWriter, r *http.Request) {
	sessionId := getCookie(r)
	user := getUserBySessionId(sessionId)
	if user == nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	if r.Method != "POST" {
		showError(w, 405, "405 Method Not Allowed")
		return
	}
	templ, err := template.ParseFiles("templates/comment.html")
	if err != nil {
		showError(w, 500, "500 Internal Server Error")
		return
	}
	postIdStr := r.FormValue("postId")
	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		showError(w, 500, "500 Internal Server Error")
		return
	}
	newCommentObject := NewCommentObject{}
	newCommentObject.User = user
	post, err := getPostById(postId)
	if err != nil {
		showError(w, 500, "500 Internal Server Error")
		return
	}
	newCommentObject.Post = post
	err = templ.Execute(w, newCommentObject)
	if err != nil {
		showError(w, 500, "500 Internal Server Error")
		return
	}
}
func commentsubmitHandler(w http.ResponseWriter, r *http.Request) {
	sessionId := getCookie(r)
	user := getUserBySessionId(sessionId)
	if user == nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	if r.Method != "POST" {
		showError(w, 405, "405 Method Not Allowed")
		return
	}
	postIdStr := r.FormValue("postId")
	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	comment := r.FormValue("comment")
	err = saveComment(user, postId, comment)
	if err != nil {
		showError(w, 500, "500 Internal Server Error")
		return
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
func setfilterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		showError(w, 405, "405 Method Not Allowed")
		return
	}
	filterCategory := r.FormValue("filterCategory")
	if !contains(filterCategories, filterCategory) {
		filterCategories = append(filterCategories, filterCategory)
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
func removefilterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		showError(w, 405, "405 Method Not Allowed")
		return
	}
	filterCategory := r.FormValue("filterCategory")
	for i, cat := range filterCategories {
		if cat == filterCategory {
			filterCategories = append(filterCategories[:i], filterCategories[i+1:]...)
			break
		}
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

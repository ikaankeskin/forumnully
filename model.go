package main

type User struct {
	Id        int
	Email     string
	Username  string
	Password  string
	SessionId string
}

type Post struct {
	Id         int
	Userid     int
	Date       int
	Content    string
	Categories []string
}

type Comment struct {
	Id      int
	Date    int
	Userid  int
	Postid  int
	Content string
}

type PostLike struct {
	Id     int
	Userid int
	Postid int
	Status int
}

type CommentLike struct {
	Id        int
	Userid    int
	Commentid int
	Status    int
}

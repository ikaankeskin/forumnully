package main

import (
	"fmt"
	"log"
)

//     _________post_likes_____________________________
//    |  id       |  userid   |  postid   |  status   |
//    |  INTEGER  |  INTEGER  |  INTEGER  |  INTEGER  |
func creratePostLikesTable() error {
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS post_likes (id INTEGER PRIMARY KEY, userid INTEGER NOT NULL, postid INTEGER NOT NULL, status INTEGER NOT NULL CHECK(status = 1 OR status = 0 OR status = -1))")
	if err != nil {
		return err
	}
	defer statement.Close()
	statement.Exec()
	return nil
}
func updatePostLikes(user *User, postId int, status int) {
	//Check if user tryes to like own post
	rows, err := db.Query("SELECT * FROM posts WHERE id = ? AND userid = ?", postId, user.Id)
	if err != nil {
		return
	}
	defer rows.Close()
	if rows.Next() {
		return
	}
	err = rows.Err()
	if err != nil {
		return
	}
	//Try to update
	statement, err := db.Prepare("UPDATE post_likes SET status = ? WHERE userid = ? AND postid = ?")
	if err != nil {
		return
	}
	defer statement.Close()
	result, err := statement.Exec(status, user.Id, postId)
	if err != nil {
		return
	}
	numbOfRows, err := result.RowsAffected()
	if err != nil {
		return
	}
	if numbOfRows == 0 {
		statement1, err := db.Prepare("INSERT INTO post_likes (userid, postid, status) VALUES (?,?,?)")
		if err != nil {
			return
		}
		defer statement1.Close()
		statement1.Exec(user.Id, postId, status)
	}
}
func printPostLikes() {
	rows, err := db.Query("SELECT * FROM post_likes")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	postLike := PostLike{}
	for rows.Next() {
		err = rows.Scan(&(postLike.Id), &(postLike.Userid), &(postLike.Postid), &(postLike.Status))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(postLike)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

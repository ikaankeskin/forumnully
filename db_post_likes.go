package main

//     _________post_likes_____________________________
//    |  id       |  userid   |  postid   |  status   |
//    |  INTEGER  |  INTEGER  |  INTEGER  |  INTEGER  |

// create post_likes table
func creratePostLikesTable() error {
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS post_likes(id INTEGER PRIMARY KEY, userid INTEGER NOT NULL, postid INTEGER NOT NULL, status INTEGER NOT NULL CHECK(status = 1 OR status = 0 OR status = -1))")
	if err != nil {
		return err
	}
	defer statement.Close()
	statement.Exec()
	return nil
}

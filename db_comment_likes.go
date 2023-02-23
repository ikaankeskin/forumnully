package main

//     _________comment_likes__________________________
//    |  id       |  userid   |  commentid  |  status  |
//    |  INTEGER  |  INTEGER  |  INTEGER    |  INTEGER |

// Create comment_likes table
func crerateCommentLikesTable() error {
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS comment_likes(id INTEGER PRIMARY KEY, userid INTEGER NOT NULL, commentid INTEGER NOT NULL, status INTEGER NOT NULL CHECK(status = 1 OR status = 0 OR status = -1))")
	if err != nil {
		return err
	}
	defer statement.Close()
	statement.Exec()
	return nil
}

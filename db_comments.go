package main

//       ________comments___________________________________________
//      |  id       |  date     |  userid   |  postid   |  content  |
//      |  INTEGER  |  INTEGER  |  INTEGER  |  INTEGER  |  TEXT     |

// Create comments table
func crerateCommentsTable() error {
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS comments(id INTEGER PRIMARY KEY, date INTEGER NOT NULL, userid INTEGER NOT NULL, postid INTEGER NOT NULL, content TEXT NOT NULL)")
	if err != nil {
		return err
	}
	defer statement.Close()
	statement.Exec()
	return nil
}

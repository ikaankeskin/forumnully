package main

//      _________posts________________________________________________
//     |  id       |  userid   |  date     |  content  |  categories  |
//     |  INTEGER  |  INTEGER  |  INTEGER  |  TEXT     |  TEXT        |

//Create posts table
func creratePostsTable() error {
	table, err := db.Prepare("CREATE TABLE IF NOT EXISTS posts(id INTEGER PRIMARY KEY, userid INTEGER NOT NULL, date INTEGER NOT NULL, content TEXT NOT NULL, categories TEXT NOT NULL)")
	if err != nil {
		return err
	}
	defer table.Close()
	table.Exec()
	return nil
}

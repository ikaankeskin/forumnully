package main

//Categories sample
//"C++", "C#", "Java", "JavaScript", "HTML", "CSS", "PHP", "Go", "Rust", "Node"})

//       __categories__
//      |  category    |
//      |  TEXT        |

// Create categories table
func crerateCategoriesTable() error {
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS categories(category TEXT NOT NULL UNIQUE)")
	if err != nil {
		return err
	}
	defer statement.Close()
	statement.Exec()
	return nil
}

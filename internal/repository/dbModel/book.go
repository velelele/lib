package dbModel

type Book struct {
	Title   string `db:"title"`
	Author  string `db:"author"`
	Body    string `db:"body"`
	Release string `db:"release"`
}

package models

type Book struct {
	ID     string `db:"id" json:"id"`
	Isbn   string `db:"isbn" json:"isbn"`
	Title  string `db:"title" json:"title"`
	Author Author `db:"author" json:"author"`
}

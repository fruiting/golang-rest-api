package controllers

import (
	"encoding/json"
	"go-rest/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetBooks - returns json of all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	results, err := models.GetDatabase().Query(
		"SELECT books.*, authors.first_name, authors.last_name FROM books " +
			"JOIN book_author ON books.id = book_author.book_id JOIN authors ON book_author.author_id = authors.id",
	)
	if err != nil {
		panic(err.Error())
	}

	var books []models.Book
	for results.Next() {
		var book models.Book
		var author models.Author

		err = results.Scan(&book.ID, &book.Isbn, &book.Title, &author.Firstname, &author.Lastname)
		if err != nil {
			panic(err.Error())
		}

		book.Author = author
		books = append(books, book)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// GetBook - returns json of specific book by ID
func GetBook(w http.ResponseWriter, r *http.Request) {
	//GET params
	params := mux.Vars(r)

	result, error := models.GetDatabase().Query(
		"SELECT books.*, authors.first_name, authors.last_name FROM books " +
			"JOIN book_author ON books.id = book_author.book_id JOIN authors ON book_author.author_id = authors.id " +
			"WHERE books.id = " + params["id"],
	)
	if error != nil {
		panic(error.Error())
	}

	var book models.Book
	var author models.Author

	result.Next()
	error = result.Scan(&book.ID, &book.Isbn, &book.Title, &author.Firstname, &author.Lastname)
	if error != nil {
		panic(error.Error())
	}

	book.Author = author
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// CreateBook - creates a book in database by json with necessary fields
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	var author models.Author
	_ = json.NewDecoder(r.Body).Decode(&book)
	_ = json.NewDecoder(r.Body).Decode(&author)

	models.GetDatabase().Query("START TRANSACTION;")
	bookResult, _ := models.GetDatabase().Exec("INSERT INTO books (isbn, title) VALUES ('" + book.Isbn + "', '" + book.Title + "');")
	bookID, _ := bookResult.LastInsertId()

	authorResult, _ := models.GetDatabase().Exec(
		"INSERT INTO authors (first_name, last_name) VALUES ('" + book.Author.Firstname + "', '" + book.Author.Lastname + "');",
	)
	authorID, _ := authorResult.LastInsertId()
	models.GetDatabase().Query(
		"INSERT INTO book_author (book_id, author_id) " +
			"VALUES (" + strconv.FormatInt(bookID, 10) + ", " + strconv.FormatInt(authorID, 10) + ");",
	)
	_, error := models.GetDatabase().Query("COMMIT;")
	if error != nil {
		panic(error.Error())
	}

	w.WriteHeader(200)
}

// UpdateBook - updates a book fields
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	//PUT params
	params := mux.Vars(r)

	var book models.Book
	var author models.Author
	_ = json.NewDecoder(r.Body).Decode(&book)
	_ = json.NewDecoder(r.Body).Decode(&author)

	models.GetDatabase().Query("START TRANSACTION;")
	models.GetDatabase().Query(
		"UPDATE books SET isbn = '" + book.Isbn + "', title = '" + book.Title + "';",
	)
	models.GetDatabase().Query(
		"UPDATE authors " +
			"JOIN book_author ON book_author.author_id = authors.id " +
			"JOIN books ON books.id = book_author.book_id " +
			"SET first_name = '" + book.Author.Firstname + "', last_name = '" + book.Author.Lastname + "' " +
			"WHERE books.id = " + params["id"] + ";",
	)
	_, error := models.GetDatabase().Query("COMMIT;")
	if error != nil {
		panic(error.Error())
	}

	w.WriteHeader(200)
}

// DeleteBook - deletes a book by ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	//DELETE params
	params := mux.Vars(r)

	_, error := models.GetDatabase().Query("DELETE FROM books WHERE id = ?", params["id"])
	if error != nil {
		panic(error.Error())
	}

	w.WriteHeader(200)
}

package models

type Book struct {
	BookName      string `json:"book_name"`
	AuthorName    string `json:"author_name"`
	PublishedYear string `json:"published_year"`
	Quantity      int    `json:"quantity"`
}

type BookName struct {
	Name string `json:"name"`
}

type Books struct {
	Books []Book
}
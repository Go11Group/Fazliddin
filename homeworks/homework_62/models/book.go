package models

type Books struct {
	BookName      string `json:"book_name"`
	AuthorName    string `json:"author_name"`
	PublishedYear string `json:"published_year"`
	Quantity      int    `json:"quantity"`
}

package models

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var Books = []Book{
	{ID: "1", Title: "Преступление и наказание", Author: "Фёдор Достоевский", Quantity: 3},
	{ID: "2", Title: "Великий Гэтсби", Author: "Ф. Скотт Фицджеральд", Quantity: 5},
	{ID: "3", Title: "Война и мир", Author: "Лев Толстой", Quantity: 6},
	{ID: "4", Title: "Мастер и Маргарита", Author: "Михаил Булгаков", Quantity: 2},
	{ID: "5", Title: "Евгений Онегин", Author: "Александр Пушкин", Quantity: 4},
}

func GetBookByID(id string) (*Book, bool) {
	for i := range Books {
		if Books[i].ID == id {
			return &Books[i], true
		}
	}
	return nil, false
}

func FilterBooksByAuthor(author string) []Book {
	var filtered []Book
	for _, b := range Books {
		if ContainsIgnoreCase(b.Author, author) {
			filtered = append(filtered, b)
		}
	}
	return filtered
}

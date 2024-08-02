package models

func SeedDB() {
	books := []Book{
		{GUID: "1234", Name: "The Great Gatsby", Author: "F. Scott Fitzgerald"},
		{GUID: "5678", Name: "1984", Author: "George Orwell"},
		{GUID: "9101", Name: "To Kill a Mockingbird", Author: "Harper Lee"},
	}

	for _, book := range books {
		DB.FirstOrCreate(&book, Book{GUID: book.GUID})
	}

	users := []User{
		{Email: "admin@example.com", Password: "password", IsAdmin: true},
		{Email: "user@example.com", Password: "password", IsAdmin: false},
	}

	for _, user := range users {
		DB.FirstOrCreate(&user, User{Email: user.Email})
	}
}

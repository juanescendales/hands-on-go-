// challenges/types-composite/begin/main.go
package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

// define an author type with a name
type author struct {
	name string
}

// define a book type with a title and author
type book struct {
	title string
	author author
}

// define a library type to track a list of books
type library struct {
	books map[string][]book
}

// define addBook to add a book to the library
func (l *library) addBook(newBook book) {
	authorBooks := l.books[newBook.author.name]
	authorBooks = append(authorBooks, newBook)
	l.books[newBook.author.name] = authorBooks
}

// define a lookupByAuthor function to find books by an author's name
func (l *library) lookupByAuthor(authorName string) []book {
	return l.books[authorName]
}

func main() {
	authorJuan := author{
		name: "Juan Esteban Cendales",
	}
	authorAndres := author{
		name: "Jaime Andres Cendales",
	}
	// create a new library
	library := library{
		books: make(map[string][]book),
	}

	// add 2 books to the library by the same author
	library.addBook(book{
		title:  "Lady Masacre",
		author: authorJuan,
	})

	library.addBook(book{
		title:  "Scorpio City",
		author: authorJuan,
	})

	// add 1 book to the library by a different author
	library.addBook(book{
		title:  "Como ser un estoico",
		author: authorAndres,
	})

	// dump the library with spew
	spew.Dump(library)

	// lookup books by known author in the library
	juanBooks := library.lookupByAuthor("Juan Esteban Cendales")
	andresBooks := library.lookupByAuthor("Jaime Andres Cendales")
	ukBooks := library.lookupByAuthor("uk")

	fmt.Printf("%#v\n", juanBooks)
	fmt.Printf("%#v\n", andresBooks)
	fmt.Printf("%#v\n", ukBooks)

	// print out the first book's title and its author's name
	fmt.Printf("Book : %v\nAuthor: %v\n", juanBooks[0].title, juanBooks[0].author.name)

}

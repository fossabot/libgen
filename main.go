package main

import "fmt"

import "github.com/binodsh/anybooks/api/libgen"

func main() {
	fmt.Println("welcome to anybooks :)")

	books := libgen.SearchBookByTitle("the power of now")

	for _, book := range books {
		fmt.Printf("%s\n%s\n%s\n%s\n\n", book.Title, book.Author, book.FileExtension, book.DownloadLink)
	}

}

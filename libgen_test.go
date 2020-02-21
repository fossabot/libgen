package libgen

import (
	"fmt"
	"testing"
)

func TestAbs(t *testing.T) {
	books := SearchBookByTitle("attitude is")

	for _, book := range books {
		fmt.Printf("Title: %s\nAuthor: %s\nExtension: %s\nDownload Link: %s\n\n", book.Title, book.Author, book.FileExtension, book.DownloadLink)
	}
}

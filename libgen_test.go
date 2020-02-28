package libgen

import (
	"fmt"
	"testing"
)

func TestAbs(t *testing.T) {
	books := SearchBookByTitle("attitude is", SortOptions{SortBy: "title", SortMode: "desc"})

	for _, book := range books {
		fmt.Printf("id:%d\nauthor:%s\ntitle:%s\n\n", book.ID, book.Author, book.Title)
	}
}

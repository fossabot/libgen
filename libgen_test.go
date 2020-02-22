package libgen

import (
	"fmt"
	"testing"
)

func TestAbs(t *testing.T) {
	books := SearchBookByTitle("attitude is", "")

	for _, book := range books {
		fmt.Println(book)
		fmt.Println()
	}
}

package libgen

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

//returns goquery.Document built from given url
func getDocument(requestURL string) (*goquery.Document, error) {
	res, err := http.Get(requestURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	return goquery.NewDocumentFromReader(res.Body)
}

func listToCSV(list []int64) string {
	csvStr := ""
	for key, id := range list {
		csvStr += fmt.Sprint(id)
		if key < len(list)-1 {
			csvStr += ", "
		}
	}
	return csvStr
}

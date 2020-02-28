package libgen

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const baseURL = "http://libgen.is/"

const searchURL = baseURL + "search.php?lg_topic=libgen&open=0&view=simple&res=100&phrase=1&column=def&"

const jsonURL = baseURL + "json.php"

const downloadURL = baseURL + "get.php"

const defaultSortBy = "def"
const defaultSortMode = "asc"

type SortOptions struct {
	SortBy   string
	SortMode string
}

//SearchBookByTitle returns the list of BookInfo which contains the search string
func SearchBookByTitle(searchStr string, sortOptions SortOptions) []BookInfo {
	sortBy := sortOptions.SortBy
	sortMode := strings.ToUpper(sortOptions.SortMode)

	// URL encode given search string
	values := url.Values{"req": {searchStr}, "sort": {sortBy}, "sortmode": {sortMode}}
	requestURL := searchURL + values.Encode()

	ids := scrapBookIdsFromSite(requestURL)
	books, err := FindBooksByIds(listToCSV(ids))

	if err != nil {
		fmt.Println(err.Error())
		return []BookInfo{}
	}

	for i, book := range books {
		books[i].DownloadLink = downloadURL + "?md5=" + book.MD5
	}

	//find books by id always returns the result sorted by ID in asc
	//so sorting the result on the basis of the ids that we get by scraping site
	var sortedBooks []BookInfo
	for _, id := range ids {
		for _, book := range books {
			if book.ID == id {
				sortedBooks = append(sortedBooks, book)
				break
			}
		}
	}

	return sortedBooks
}

func FindBooksByIds(ids string) ([]BookInfo, error) {
	fmt.Println("ids: " + ids)
	params := url.Values{"ids": {ids}, "fields": {"*"}}

	requestURL := jsonURL + "?" + params.Encode()

	res, err := http.Get(requestURL)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	var books []BookInfo
	err = json.NewDecoder(res.Body).Decode(&books)

	if err != nil {
		return books, fmt.Errorf("erorr while decoding json")
	}

	return books, nil
}

//scrapBookIdsFromSite loads the page of given url (libgen.is) and gets all the ids from the table
func scrapBookIdsFromSite(requestURL string) []int64 {
	doc, err := getDocument(requestURL)
	if err != nil {
		log.Fatal(err)
	}

	var ids []int64
	doc.Find(".c > tbody:nth-child(1) > tr").Each(func(i int, s *goquery.Selection) {
		if i != 0 {
			//reach to id column
			id := s.Find("td:nth-child(1)").Text()
			idInt, _ := strconv.ParseInt(id, 10, 64)
			ids = append(ids, idInt)
		}
	})

	return ids
}

//returns goquery.Document build from given url
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

package libgen

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const baseURL = "http://libgen.is/"

const searchURL = baseURL + "search.php?lg_topic=libgen&open=0&view=simple&res=100&phrase=1&column=def&"

const defaultFields = "id,author,title,md5,year,extension,filesize"

const jsonURL = baseURL + "json.php"

const downloadURL = baseURL + "get.php"

//SearchBookByTitle returns the list of BookInfo which contains the search string
func SearchBookByTitle(searchStr string, fields string) []BookInfo {
	// URL encode given search string
	value := url.Values{"req": {searchStr}}
	requestURL := searchURL + value.Encode()

	doc, err := getDocument(requestURL)
	if err != nil {
		log.Fatal(err)
	}

	var ids string
	doc.Find(".c > tbody:nth-child(1) > tr").Each(func(i int, s *goquery.Selection) {
		//reach to id column
		id := s.Find("td:nth-child(1)").Text()
		//make csv of ids
		ids += id + ","
	})

	if ids == "" {
		fmt.Println("no books found")
		return []BookInfo{}
	}

	books, err := FindBooksByIds(ids, fields)

	if err != nil {
		return []BookInfo{}
	}

	for i, book := range books {
		books[i].DownloadLink = downloadURL + "?md5=" + book.MD5
	}

	return books
}

func FindBooksByIds(ids string, fields string) ([]BookInfo, error) {

	if strings.TrimSpace(fields) == "" {
		fields = defaultFields
	}

	params := url.Values{"ids": {ids}, "fields": {fields}}

	requestURL := jsonURL + "?" + params.Encode()
	fmt.Println("request url by ids " + requestURL)

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

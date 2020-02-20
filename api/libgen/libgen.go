package libgen

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

const baseURL = "http://libgen.is/"

const searchURL = baseURL + "search.php?lg_topic=libgen&open=0&view=simple&res=100&phrase=1&column=def&"

const jsonURL = baseURL + "json.php?fields=id,author,title,md5,year,extension,filesize&"

const downloadURL = baseURL + "get.php?md5="

func SearchBookByTitle(title string) []BookInfo {
	doc, err := getDocument(title)
	if err != nil {
		log.Fatal(err)
	}

	var ids string

	doc.Find(".c > tbody:nth-child(1) > tr").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title

		id := s.Find("td:nth-child(1)").Text()

		ids += id + ","
	})

	if ids == "" {
		fmt.Println("no books found")
		return []BookInfo{}
	}

	books, err := findBooksByIds(ids)

	if err != nil {
		return []BookInfo{}
	}

	for i, book := range books {
		books[i].DownloadLink = downloadURL + book.MD5
	}

	return books
}

func findBooksByIds(ids string) ([]BookInfo, error) {
	requestURL := jsonURL + "ids=" + ids
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

func getDocument(title string) (*goquery.Document, error) {
	value := url.Values{"req": {title}}
	requestURL := searchURL + value.Encode()
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

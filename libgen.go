package libgen

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const baseURL = "http://libgen.is/"

const mirror2BaseURL = "http://libgen.lc/"

const searchURL = baseURL + "search.php?lg_topic=libgen&open=0&view=simple&res=100&phrase=1&column=def&"

const jsonURL = baseURL + "json.php"

const downloadURL = mirror2BaseURL + "ads.php"

type SortOptions struct {
	SortBy   string
	SortMode string
}

//SearchBookByTitle returns the list of BookInfo which contains the search string
func SearchBookByTitle(searchStr string, sortOptions SortOptions) ([]BookInfo, error) {
	sortBy := sortOptions.SortBy
	sortMode := strings.ToUpper(sortOptions.SortMode)

	// URL encode given search string
	values := url.Values{"req": {searchStr}, "sort": {sortBy}, "sortmode": {sortMode}}
	requestURL := searchURL + values.Encode()

	ids := scrapBookIdsFromSite(requestURL)
	books, err := FindBooksByIds(ids)

	if err != nil {
		return []BookInfo{}, err
	}

	for i, book := range books {
		books[i].DownloadPageUrl = downloadURL + "?md5=" + book.MD5
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

	return sortedBooks, nil
}

func FindBooksByIds(ids []int64) ([]BookInfo, error) {
	if len(ids) == 0 {
		return []BookInfo{}, nil
	}

	idsCSV := listToCSV(ids)

	fmt.Println("ids: " + idsCSV)
	params := url.Values{"ids": {idsCSV}, "fields": {"*"}}

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

func GetDownloadInfo(bookID int64) (DownloadInfo, error) {
	books, err := FindBooksByIds([]int64{bookID})

	if err != nil {
		return DownloadInfo{}, err
	}

	if len(books) == 0 {
		return DownloadInfo{}, errors.New("book not found")
	}

	book := books[0]

	downloadPageURL := downloadURL + "?md5=" + book.MD5

	doc, err := getDocument(downloadPageURL)
	if err != nil {
		return DownloadInfo{}, nil
	}

	//select the GET button of the page
	link, _ := doc.Find("#main > tbody:nth-child(1) > tr:nth-child(1) > td:nth-child(2) > a:nth-child(1)").Attr("href")
	return DownloadInfo{
		ID:              bookID,
		Title:           book.Title,
		DownloadPageURL: downloadPageURL,
		DowloadLink:     link,
	}, nil
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

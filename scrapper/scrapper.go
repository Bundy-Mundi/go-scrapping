package scrapper

import (
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Data type
type Data struct {
	title  string
	rank   string
	artist string
	date   time.Time
}

// Scrapper -
type Scrapper struct {
	urls []string
}

// NewScrapper - Scrapper constructor
func NewScrapper(urls []string) *Scrapper {
	return &Scrapper{urls: urls}
}

// Scrape - ...
func (s *Scrapper) Scrape() map[string][]Data {
	result := make(map[string][]Data)
	if s.urls != nil {
		for _, url := range s.urls {
			r := fetchURL(url)
			result[url] = r
		}
	}

	return result
}

func fetchURL(url string) []Data {
	result := []Data{}
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Find the review items
	doc.Find("li.chart-list__element").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		button := s.Find("button.chart-element__wrapper")
		rank := button.
			Find("span.chart-element__rank").
			Find("span.chart-element__rank__number").Text()
		title := button.Find("span.chart-element__information").Find("span.chart-element__information__song").Text()
		artist := button.Find("span.chart-element__information").Find("span.chart-element__information__artist").Text()
		result = append(result, Data{title: title, rank: rank, artist: artist})
	})

	return result
}

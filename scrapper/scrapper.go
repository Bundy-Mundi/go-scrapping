package scrapper

import (
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Data type
type Data struct {
	title  string    `json: "title"`
	rank   string    `json: "rank"`
	artist string    `json: "artist"`
	date   time.Time `json: "created_at"`
}

// Scrapper -
type Scrapper struct {
	urls []string // will be used later
}

// NewScrapper - Scrapper constructor
func NewScrapper(urls []string) *Scrapper {
	return &Scrapper{urls: urls}
}

// ScrapeBillboard - Scrapper for billboard chart
func (s *Scrapper) ScrapeBillboard(doc *goquery.Document) []*Data {
	var result []*Data
	var d *Data

	// Find the review items
	doc.Find("li.chart-list__element").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		button := s.Find("button.chart-element__wrapper")
		rank := button.
			Find("span.chart-element__rank").
			Find("span.chart-element__rank__number").Text()
		title := button.Find("span.chart-element__information").Find("span.chart-element__information__song").Text()
		artist := button.Find("span.chart-element__information").Find("span.chart-element__information__artist").Text()
		d = &Data{title: title, rank: rank, artist: artist}
		result = append(result, d)
	})

	return result
}

// FetchURL - Return html doc
func (s *Scrapper) FetchURL(url string) *goquery.Document {

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
	return doc
}

// SaveFileTXT - Save []*Data as txt file
func (s *Scrapper) SaveFileTXT() {

}

// SaveFileCSV - Save []*Data as csv file
func (s *Scrapper) SaveFileCSV() {

}

// SaveFileJSON - Save []*Data as json file
func (s *Scrapper) SaveFileJSON() {

}

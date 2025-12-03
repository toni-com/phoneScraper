package scraper

import (
	"github.com/gocolly/colly"
)

// Scrape TODO: currently overwrites every result it finds -> only return last result
func Scrape(url string, sel string) (string, error) {
	c := colly.NewCollector()
	var result string
	c.OnHTML(sel, func(e *colly.HTMLElement) {
		result = e.Text
	})
	err := c.Visit(url)
	if err != nil {
		return "", err
	}
	return result, nil
}

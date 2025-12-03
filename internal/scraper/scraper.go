package scraper

import (
	"github.com/gocolly/colly"
)

func Scrape(url string, sel string) (string, error) {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"),
	)
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

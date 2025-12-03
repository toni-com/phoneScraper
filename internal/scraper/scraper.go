package scraper

import (
	"os"
	"strings"

	"github.com/gocolly/colly"
)

// Scrape TODO: currently overwrites every result it finds -> only return last result
func Scrape(url string, sel string) (string, error) {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"),
	)
	var result string

	c.OnHTML(sel, func(e *colly.HTMLElement) {
		result = e.Text
	})

	c.OnResponse(func(r *colly.Response) {
		// debug purposes
		if strings.Contains(r.Request.URL.String(), "mediamarkt") {
			os.WriteFile("debug_mediamarkt.html", r.Body, 0644)
		}
	})
	err := c.Visit(url)
	if err != nil {
		return "", err
	}
	return result, nil
}

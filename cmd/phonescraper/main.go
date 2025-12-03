package main

import (
	"PhoneScraper/internal/config"
	"PhoneScraper/internal/scraper"
	"fmt"
	"log"
	"sync"
)

func main() {
	// Load Config
	items, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}

	//Setup Concurrency Control
	var wg sync.WaitGroup

	fmt.Printf("Starting sniper on %d items...\n", len(items))

	//Loop over items
	for _, item := range items {
		wg.Add(1)

		// Goroutine
		go func(i config.ItemConfig) {
			defer wg.Done()
			scrapeResult, err := scraper.Scrape(i.URL, i.Selector)
			if err != nil {
				log.Printf("Failed to scrape price %s: %v", i.Name, err)
				return
			}
			priceClean, err := scraper.ParsePrice(scrapeResult)
			if err != nil {
				log.Printf("Failed to parse price %s: %v", i.Name, err)
				return
			}
			if priceClean < i.Threshold {
				fmt.Printf("Alert: {%s} on sale for {%.2f}", i.Name, priceClean)
			}
		}(item)
	}
	wg.Wait()
	fmt.Println("All checks complete.")
}

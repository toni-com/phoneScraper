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
	items, err := config.LoadConfig("config2.json")
	if err != nil {
		log.Fatal(err)
	}

	//Setup Concurrency Control
	var wg sync.WaitGroup
	jobs := make(chan config.ItemConfig, len(items))
	const workerLimit int = 3

	// start worker
	for i := 0; i < workerLimit; i++ {
		wg.Add(1)
		// goroutine
		go func(j chan config.ItemConfig) {
			defer wg.Done()
			// loop over job
			for item := range j {
				scrapeResult, err := scraper.Scrape(item.URL, item.Selector)
				if err != nil {
					log.Printf("Failed to scrape price %s: %v\n", item.Name, err)
					continue
				}
				priceClean, err := scraper.ParsePrice(scrapeResult)
				if err != nil {
					log.Printf("Failed to parse price %s: %v\n", item.Name, err)
					continue
				}
				if priceClean < item.Threshold {
					fmt.Printf("Alert: {%s} on sale for {%.2f}\n", item.Name, priceClean)
				}
			}
		}(jobs)
	}

	fmt.Printf("Starting sniper on %d items...\n", len(items))
	for _, item := range items {
		jobs <- item
	}
	close(jobs)

	wg.Wait()
	fmt.Println("All checks complete.")
}

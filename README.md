# Smartphone Sniper 
*A lightweight, concurrent web scraper written in Go.*

I built this project to learn idiomatic Go patterns, specifically focusing on concurrency pipelines and real-world HTML parsing.

It reads a list of target URLs from a config file, scrapes the current price, and triggers an alert if the item is below a set threshold.

---

## Demo
| Scraping Results |
| :---: |
| <img width="618" height="172" alt="phonescraper" src="https://github.com/user-attachments/assets/64026992-d139-4a72-aa39-867b93e8f836" /> |

---

## Tech Stack

- **Language:** Go 
- **Scraping Framework:** Colly  

---

## Key Learnings


### **Concurrency Patterns**
Moved from basic Goroutines to a Worker Pool pattern to manage resources and prevent rate-limiting.

### **TDD**
Practiced Test-Driven Development by mocking HTML responses to test the scraping and cleaning logic before hitting real sites.

### **Robustness**
Handled edge cases like German number formatting, 403/404 errors, and empty DOM elements without crashing worker threads.


---

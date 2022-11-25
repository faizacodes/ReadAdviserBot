package main

import (
	"faiza_bot/bots"
	"faiza_bot/scraping"
	"time"
)

func main() {
	go scraping.ScrapingFunc()
	time.Sleep(3 * time.Second)
	bots.BotFunc()
}

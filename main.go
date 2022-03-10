package main

import "github.com/plusbeauxjours/jobScrapper-GO/scrapper"

func main() {
	scrapper.Scrape("term")
}

// go env -w GO111MODULE=on
// go get github.com/labstack/echo/v4

// go env -w GO111MODULE=on
// go get github.com/PuerkitoBio/goquery

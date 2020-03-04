package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		// Only visit domains allowed
		colly.AllowedDomains("wiki.xandr.com", "wiki.appnexus.com"),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print out the link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)

		// Visit link found
		// Since AllowedDomains were set, only visiting expected links
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Log out each site being visited
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Now Visiting", r.URL.String())
	})

	c.Visit("https://wiki.xandr.com/")

}

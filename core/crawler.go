package core

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

var UrlImage string
var UrlVideo string

func CrawlWeb(input string) {
	c := colly.NewCollector()

	// Find and visit all links
	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	e.Request.Visit(e.Attr("href"))
	// })
	c.OnHTML(`html`, func(e *colly.HTMLElement) {
		fmt.Println("success get element")
		fmt.Println(e.ChildAttr(`meta[property="og:image"]`, "content"))
		UrlImage = e.ChildAttr(`meta[property="og:image"]`, "content")
		UrlVideo = e.ChildAttr(`meta[property="og:video:secure_url"]`, "content")

	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(input)
}

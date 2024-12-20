package production

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1085", "美国广播公司", "https://abcnews.go.com/")

	engine.SetStartingURLs([]string{
		"https://abcnews.go.com/US",
		"https://abcnews.go.com/Business",
		"https://abcnews.go.com/Entertainment",
		"https://abcnews.go.com/Politics",
		"https://abcnews.go.com/Technology",
		"https://abcnews.go.com/International",
	})
	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         true,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnHTML(".ContentList__Item > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".ContentRoll__Headline > h2 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		fileURL := element.Attr("href")
		if strings.Contains(fileURL, "video") {
			url, err := element.Request.URL.Parse(element.Attr("href"))
			if err != nil {
				crawlers.Sugar.Error(err.Error())
				return
			}
			ctx.Video = append(ctx.Video, url.String())
			ctx.PageType = crawlers.Report
		} else {
			ctx.Visit(element.Attr("href"), crawlers.News)
		}
	})
}

package parser

import (
	"concurrencyParser/models"

	"github.com/gocolly/colly"
)

func ScrapeRia(URL string) ([]models.Response, error) {
	c := colly.NewCollector(colly.Async(true))
	var res []models.Response
	c.OnHTML("a.cell-list__item-link", func(e *colly.HTMLElement) {
		r := models.Response{Header: e.Text, Link: e.Attr("href")}
		res = append(res, r)
	})
	c.Visit(URL)
	c.Wait()
	return res, nil
}

func ScrapeRbk(URL string) ([]models.Response, error) {
	c := colly.NewCollector(colly.Async(true))
	var res []models.Response
	c.OnHTML("a.news-feed__item", func(e *colly.HTMLElement) {
		r := models.Response{Header: e.Text, Link: e.Attr("href")}
		res = append(res, r)
	})
	c.Visit(URL)
	c.Wait()
	return res, nil
}

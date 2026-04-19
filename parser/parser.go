package parser

import (
	"concurrencyParser/models"
	"concurrencyParser/storage"
	"database/sql"
	"fmt"

	"github.com/gocolly/colly"
)

func ScrapeRia(URL string, db *sql.DB) ([]models.Response, error) {
	c := colly.NewCollector(colly.Async(true))
	var res []models.Response
	c.OnHTML("a.cell-list__item-link", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		r := models.Response{Header: e.Text, Link: link}
		resp, err := storage.Insert(db, r)
		if err != nil {
			fmt.Println(err)
		}
		res = append(res, resp)
	})
	c.Visit(URL)
	c.Wait()
	return res, nil
}

func ScrapeRbk(URL string, db *sql.DB) ([]models.Response, error) {
	c := colly.NewCollector(colly.Async(true))
	var res []models.Response
	c.OnHTML("a.card-mini", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		r := models.Response{Header: e.Text, Link: link}
		resp, err := storage.Insert(db, r)
		if err != nil {
			fmt.Println(err)
		}
		res = append(res, resp)
	})
	c.Visit(URL)
	c.Wait()
	return res, nil
}

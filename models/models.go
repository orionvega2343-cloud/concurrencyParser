package models

type Response struct {
	ID     int    `json:"id"`
	Header string `json:"header"`
	Link   string `json:"link"`
}

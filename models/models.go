package models

type Response struct {
	ID     int    `json:"id"`
	Header string `json:"header"`
	Text   string `json:"text"`
}

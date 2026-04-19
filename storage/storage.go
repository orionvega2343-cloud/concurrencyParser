package storage

import (
	"concurrencyParser/models"
	"database/sql"

	_ "github.com/lib/pq"
)

func NewDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func CreateTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS responses (
    						  ID SERIAL PRIMARY KEY,
    						  header TEXT NOT NULL,
    						  link TEXT NOT NULL	
)`)
	if err != nil {
		return err
	}
	return nil

}

func Insert(db *sql.DB, r models.Response) (models.Response, error) {
	err := db.QueryRow("INSERT INTO responses (header, link) VALUES ($1, $2) RETURNING id", r.Header, r.Link).Scan(&r.ID)
	if err != nil {
		return models.Response{}, err
	}
	return r, nil
}

func GetAll(db *sql.DB) ([]models.Response, error) {
	rows, err := db.Query("SELECT id, header, link FROM responses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []models.Response
	for rows.Next() {
		var resp models.Response
		err := rows.Scan(&resp.ID, &resp.Header, &resp.Link)
		if err != nil {
			return nil, err
		}
		res = append(res, resp)
	}
	return res, nil
}

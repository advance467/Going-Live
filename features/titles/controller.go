package titles

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/advance467/Going-Live/database"
	"github.com/go-chi/chi/v5"
)

type TitleItem struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var (
	id          int
	name        string
	description string
)

func GetTitle(w http.ResponseWriter, r *http.Request) {

	titleID := chi.URLParam(r, "id")

	var statement string = `SELECT * FROM titles where id = $1 LIMIT 1`
	row := database.Client.QueryRow(statement, titleID)

	err := row.Scan(&id, &name, &description)

	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}

	if err == sql.ErrNoRows {
		w.WriteHeader(404)
		message, _ := json.Marshal("Resource Not Found")
		w.Write(message)
	} else {

		item := TitleItem{
			ID:          id,
			Name:        name,
			Description: description,
		}
		data, _ := json.Marshal(item)
		w.Write(data)
	}

}

func getTitles(w http.ResponseWriter, r *http.Request) {

	var statement string = `SELECT * FROM titles`
	rows, err := database.Client.Query(statement)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	results := []TitleItem{}

	for rows.Next() {

		if err := rows.Scan(&id, &name, &description); err != nil {
			log.Fatal(err)
		}

		result := TitleItem{
			ID:          id,
			Name:        name,
			Description: description,
		}

		results = append(results, result)
	}

	data, _ := json.Marshal(results)
	w.Write(data)

}

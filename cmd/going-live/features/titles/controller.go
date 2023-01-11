package titles

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/advance467/Going-Live/internal/database"
)

type TitleItem struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func getTitle(w http.ResponseWriter, r *http.Request) {

	titleID := chi.URLParam(r, "id")
	stmt := "SELECT * FROM titles where id = $1 LIMIT 1"
	row := database.GetDB().QueryRow(stmt, titleID)

	var item TitleItem
	err := row.Scan(&item.ID, &item.Name, &item.Description)

	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}

	if err == sql.ErrNoRows {
		w.WriteHeader(404)
		message, _ := json.Marshal("Resource Not Found")
		w.Write(message)
		return
	}

	data, _ := json.Marshal(item)
	w.Write(data)
}

func getTitles(w http.ResponseWriter, r *http.Request) {

	stmt := "SELECT * FROM titles"
	rows, err := database.GetDB().Query(stmt)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	items := []TitleItem{}

	for rows.Next() {
		var item TitleItem
		err := rows.Scan(&item.ID, &item.Name, &item.Description)

		if err != nil {
			log.Fatal(err)
		}

		items = append(items, item)
	}

	data, _ := json.Marshal(items)
	w.Write(data)
}

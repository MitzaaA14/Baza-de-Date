package pages

import (
	"fmt"
	"log"
	"net/http"
	"robby/db"
	"strings"
)

const (
	indexSQL = `select  t.id_turneu, t.nume_turneu, t.data_inceput, count(c.id_concert) as numar_concerte  from turnee t
		right join concerte c on c.id_turneu = t.id_turneu
		group by  t.id_turneu, t.nume_turneu,  t.data_inceput
	`
)

type turneu struct {
	IdTurneu      int
	Nume          string
	DataInceput   string
	NumarConcerte int
}

func HandlerIndexfunc(w http.ResponseWriter, r *http.Request) {
	log.Printf("request index %s", r.RequestURI)

	orderIndexSQL := indexSQL

	orderBy := r.URL.Query().Get("order")
	if orderBy != "" {
		oa := strings.Split(orderBy, ":")

		orderIndexSQL += fmt.Sprintf(" order by %s %s", oa[0], oa[1])
	}

	rows, err := db.DB.Query(orderIndexSQL)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "%s", err)
		return
	}

	var turnee []turneu

	for rows.Next() {
		var t turneu

		if err := rows.Scan(&t.IdTurneu, &t.Nume, &t.DataInceput, &t.NumarConcerte); err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "%s", err)
			return
		}

		turnee = append(turnee, t)
	}

	Render(w, r, "index.html", map[string]any{"Data": turnee}, orderIndexSQL)
}

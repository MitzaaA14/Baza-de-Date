package pages

import (
	"fmt"
	"log"
	"net/http"
	"robby/db"
	"strconv"
	"strings"
)

var turneuSQL = `
	select 
		c.id_concert,
		c.tara,
		c.oras,
		c.locatie,
		c.data_concert,
		c.capacitate,
		count(b.id_bilet) as bilete_vandute
	from concerte c
	left join bilete b on b.id_concert = c.id_concert
	where c.id_turneu = ?
	group by 
		c.id_concert,
		c.tara,
		c.oras,
		c.locatie,
		c.data_concert,
		c.capacitate
`
var getNextConcertIdSQL = `select max(id_concert)+1 from concerte;`

var insetConcertSQL = `
insert into 
	concerte(id_concert, id_turneu, tara, oras, locatie, data_concert, capacitate) 
	values(?, ?, ? , ?, ?, STR_TO_DATE(?, '%Y-%m-%d') , ?)
`

var updateConcertSQL = `
update concerte
set 
    id_turneu = ?, 
    tara = ?, 
    oras = ?, 
    locatie = ?, 
    data_concert = STR_TO_DATE(?, '%Y-%m-%d'), 
    capacitate = ?
where 
    id_concert = ?;
`

var getFansSQL = `select id_fan, concat(nume_fan, ' ', prenume_fan, ' &lt;', email, '&gt;') as fan from comunitate`

var getNextBiletIdSQL = `select max(id_bilet)+1 from bilete;`

var insetFanBiletSQL = `
insert into 
	bilete(id_bilet, id_concert, id_fan, tip_bilet, pret) 
	values(?, ?, ?, ?, ?)
`

var deleteConcertSQL = `
	delete from concerte where id_concert=? and id_turneu=?
`

type concert struct {
	IdConcert     int
	Tara          string
	Oras          string
	Locatie       string
	DataConcert   string
	Capacitate    int
	BileteVandute int
}

type fan struct {
	IdFan int
	Fan   string
}

func HandlerTurneufunc(w http.ResponseWriter, r *http.Request) {
	log.Printf("request turneu %s", r.RequestURI)

	allSQLs := ""

	if r.URL.Query().Get("action") == "delete" {
		id_turneu, _ := strconv.Atoi(r.URL.Query().Get("id"))
		concert_id, _ := strconv.Atoi(r.URL.Query().Get("concert_id"))

		params := []interface{}{
			concert_id,
			id_turneu,
		}

		_, err := db.DB.Exec(deleteConcertSQL, params...)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "%s\n%s", db.GetFinalQueryToPrint(deleteConcertSQL, params), err)
			return
		}

		allSQLs += db.GetFinalQueryToPrint(deleteConcertSQL, params) + "\n"
	}

	if r.URL.Query().Get("exec") == "edit" {
		id_turneu, _ := strconv.Atoi(r.URL.Query().Get("id"))
		if r.URL.Query().Get("concert_id") != "" {
			q := r.URL.Query()
			params := []interface{}{
				id_turneu,
				strings.TrimSpace(q.Get("tara")), strings.TrimSpace(q.Get("oras")),
				strings.TrimSpace(q.Get("locatie")), strings.TrimSpace(q.Get("data_concert")), strings.TrimSpace(q.Get("capacitate")),
				strings.TrimSpace(q.Get("concert_id")),
			}
			_, err := db.DB.Exec(updateConcertSQL, params...)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "%s\n%s", db.GetFinalQueryToPrint(updateConcertSQL, params), err)
				return
			}

			allSQLs += db.GetFinalQueryToPrint(updateConcertSQL, params) + "\n"
		} else {
			var id_concert int
			row := db.DB.QueryRow(getNextConcertIdSQL)
			if err := row.Scan(&id_concert); err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "%s", err)
				return
			}

			allSQLs += getNextConcertIdSQL + "\n"
			q := r.URL.Query()
			params := []interface{}{
				id_concert, id_turneu,
				strings.TrimSpace(q.Get("tara")), strings.TrimSpace(q.Get("oras")),
				strings.TrimSpace(q.Get("locatie")), strings.TrimSpace(q.Get("data_concert")), strings.TrimSpace(q.Get("capacitate")),
			}
			_, err := db.DB.Exec(insetConcertSQL, params...)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "%s\n%s", db.GetFinalQueryToPrint(insetConcertSQL, params), err)
				return
			}

			allSQLs += db.GetFinalQueryToPrint(insetConcertSQL, params) + "\n"
		}
	}

	if r.URL.Query().Get("exec") == "vinde_bilete" {
		var bilet_id int
		row := db.DB.QueryRow(getNextBiletIdSQL)
		if err := row.Scan(&bilet_id); err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "%s", err)
			return
		}

		q := r.URL.Query()

		allSQLs += getNextBiletIdSQL + "\n"

		var id_fan any = strings.TrimSpace(q.Get("id_fan"))
		if id_fan == "none" {
			id_fan = nil
		}

		params := []interface{}{
			bilet_id, strings.TrimSpace(q.Get("concert_id")),
			id_fan, strings.TrimSpace(q.Get("tip_bilet")),
			strings.TrimSpace(q.Get("pret_bilet")),
		}
		_, err := db.DB.Exec(insetFanBiletSQL, params...)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "%s\n%s", db.GetFinalQueryToPrint(insetFanBiletSQL, params), err)
			return
		}

		allSQLs += db.GetFinalQueryToPrint(insetFanBiletSQL, params) + "\n"
	}

	var concerte []concert
	orderTurneuSQL := turneuSQL

	filter := r.URL.Query().Get("filter_count_bilete")
	if filter != "" {
		orderTurneuSQL += fmt.Sprintf(" having bilete_vandute %s", filter)
	}

	orderBy := r.URL.Query().Get("order")
	if orderBy != "" {
		oa := strings.Split(orderBy, ":")

		orderTurneuSQL += fmt.Sprintf(" order by %s %s", oa[0], oa[1])
	}

	idTurneu := r.URL.Query().Get("id")

	rows, err := db.DB.Query(orderTurneuSQL, idTurneu)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "%s", err)
		return
	}

	allSQLs += db.GetFinalQueryToPrint(orderTurneuSQL, []interface{}{idTurneu}) + "\n"

	for rows.Next() {
		var t concert

		if err := rows.Scan(&t.IdConcert, &t.Tara, &t.Oras, &t.Locatie, &t.DataConcert, &t.Capacitate, &t.BileteVandute); err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "%s", err)
			return
		}

		concerte = append(concerte, t)
	}

	var edit concert

	if r.URL.Query().Get("action") == "edit" {
		for _, c := range concerte {
			if fmt.Sprintf("%d", c.IdConcert) == r.URL.Query().Get("concert_id") {
				edit = c
				break
			}
		}
	}

	var fani []fan

	if r.URL.Query().Get("action") == "vinde_bilet" {
		rows, err := db.DB.Query(getFansSQL)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "%s", err)
			return
		}

		allSQLs += getFansSQL + "\n"

		for rows.Next() {
			var f fan

			if err := rows.Scan(&f.IdFan, &f.Fan); err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "%s", err)
				return
			}

			fani = append(fani, f)
		}
	}

	Render(w, r, "turneu.html", map[string]any{"Data": concerte, "Edit": edit, "Fani": fani}, allSQLs)
}

package pages

import (
	"fmt"
	"log"
	"net/http"
	"robby/db"
	"strings"
)

var cautareSQL = `
select 
	t.id_turneu, t.nume_turneu,
	c.tara, c.oras, c.locatie, c.data_concert,
    p.titlu, p.durata, p.compozitor, p.link_piesa,
    a.titlu, a.gen, a.casa_de_discuri
from lista_piese_concert lp
join concerte c on lp.id_concert = c.id_concert
join turnee t on t.id_turneu = c.id_turneu
join piese p on lp.id_piesa = p.id_piesa
join album a on p.id_album = a.id_album
where 
	t.nume_turneu like ? or
	c.tara like ? or 
    c.oras like ? or 
	c.locatie like ? or 
	c.data_concert like ? or 
    p.titlu like ? or 
	p.durata like ? or 
	p.compozitor like ? or 
	p.link_piesa  like ? or
    a.titlu like ? or 
	a.gen like ? or 
	a.casa_de_discuri like ?
`

type Cautare struct {
	IdTurneu           int
	NumeTurneu         string
	Tara               string
	Oras               string
	Locatie            string
	DataConcert        string
	PiesaTitlu         string
	PiesaDurata        string
	PiesaCompozitor    string
	PiesaLink          *string
	AlbumTitlu         string
	AlbumGen           string
	AlbumCasaDeDiscuri string
}

func HandlerCautarefunc(w http.ResponseWriter, r *http.Request) {
	log.Printf("request cautare %s", r.RequestURI)

	cautare := r.URL.Query().Get("cautare")
	if cautare == "" {
		Render(w, r, "cautare.html", nil, "")
		return
	}

	orderCautareSQL := cautareSQL

	orderBy := r.URL.Query().Get("order")
	if orderBy != "" {
		oa := strings.Split(orderBy, ":")

		orderCautareSQL += fmt.Sprintf(" order by %s %s", oa[0], oa[1])
	}

	cautareArr := make([]any, 0)
	for i := 0; i < 12; i++ {
		cautareArr = append(cautareArr, fmt.Sprintf("%%%s%%", cautare))
	}

	rows, err := db.DB.Query(orderCautareSQL, cautareArr...)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "%s\n%s", db.GetFinalQueryToPrint(orderCautareSQL, cautareArr), err)
		return
	}

	var rezultate []Cautare

	for rows.Next() {
		var c Cautare

		if err := rows.Scan(
			&c.IdTurneu,
			&c.NumeTurneu,
			&c.Tara,
			&c.Oras,
			&c.Locatie,
			&c.DataConcert,
			&c.PiesaTitlu,
			&c.PiesaDurata,
			&c.PiesaCompozitor,
			&c.PiesaLink,
			&c.AlbumTitlu,
			&c.AlbumGen,
			&c.AlbumCasaDeDiscuri,
		); err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "%s", err)
			return
		}

		rezultate = append(rezultate, c)
	}

	sqlPrint := db.GetFinalQueryToPrint(orderCautareSQL, cautareArr)

	Render(w, r, "cautare.html", map[string]any{"Data": rezultate}, sqlPrint)
}

package pages

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
)

var Templates *template.Template

var funcMap = template.FuncMap{
	"paramGetToggleOrder": func(orderByName string, r *http.Request) string {
		order := r.URL.Query().Get("order")
		if order == "" {
			return upsetRequestPath(r, "order", fmt.Sprintf("%s:ASC", orderByName))
		}

		orderSplit := strings.Split(order, ":")
		name := orderSplit[0]
		by := orderSplit[1]

		if name != orderByName {
			return upsetRequestPath(r, "order", fmt.Sprintf("%s:ASC", orderByName))
		}

		if by == "" || by == "DESC" {
			return upsetRequestPath(r, "order", fmt.Sprintf("%s:ASC", orderByName))
		}

		return upsetRequestPath(r, "order", fmt.Sprintf("%s:DESC", orderByName))
	},

	"getQueryParam": func(name string, r *http.Request) string {
		return r.URL.Query().Get(name)
	},
}

func upsetRequestPath(r *http.Request, name, value string) string {
	q := r.URL.Query()

	q.Set(name, value)
	return fmt.Sprintf("%s?%s", r.URL.Path, q.Encode())
}

func InitTemplates() error {
	tmpl, err := template.New("").Funcs(funcMap).ParseGlob("pages/templates/*")
	if err != nil {
		return err
	}

	Templates = tmpl

	log.Printf("Templates inited")

	return nil
}

func Render(w http.ResponseWriter, r *http.Request, page string, data map[string]any, query string) {
	pageData := map[string]interface{}{
		"Request": r,
		"Query":   query,
	}

	for k, v := range data {
		pageData[k] = v
	}

	if err := Templates.ExecuteTemplate(w, page, pageData); err != nil {
		log.Printf("ERROR: render index: %s", err)
	}
}

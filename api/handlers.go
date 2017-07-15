package api

import (
	"html/template"
	"log"
	"net/http"
)

type Person struct {
	UserName string
}

func (m *manager) generateHandlerGetStat() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var a = template.Must(template.ParseFiles("assets/statistic.html"))
		results := m.dbs.OneMonthStatistic(1, 07, 2017)
		log.Println(results)
		a.Execute(w, results)
	}
}

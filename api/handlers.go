package api

import (
	"html/template"
	"log"
	"net/http"

	"github.com/wastebot/dbs"
)

type Result struct {
	CurMonth []dbs.Result
	YearStat []dbs.OneYearStat
}

func (m *manager) generateHandlerGetStat() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var a = template.Must(template.ParseFiles("assets/statistic.html"))
		results := m.dbs.OneMonthStatistic(1, 07, 2017)
		results2 := m.dbs.OneYearStatistic(1, 2017)
		result := Result{
			CurMonth: results,
			YearStat: results2,
		}
		log.Println(result)
		a.Execute(w, result)
	}
}

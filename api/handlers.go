package api

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wastebot/dbs"
)

var statTemplate = template.Must(template.ParseFiles("assets/statistic.html"))

type Result struct {
	CurMonth []dbs.Result
	YearStat []dbs.OneYearStat
}

func (m *manager) generateHandlerGetStat() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		urlVariables := mux.Vars(r)
		uuid := urlVariables["key"]

		// Ищем пользователя по uuid
		user, err := m.dbs.FindUserByUUID(uuid)
		if err != nil {
			// TODO: отправить специальный темплейт (в идеале 2: что то пошло не так и 404)
			log.Println(err)
			w.Write([]byte("Not found"))
			return
		}

		results := m.dbs.OneMonthStatistic(user.ID, 07, 2017)
		results2 := m.dbs.OneYearStatistic(user.ID, 2017)
		result := Result{
			CurMonth: results,
			YearStat: results2,
		}
		log.Println(result)
		statTemplate.Execute(w, result)
	}
}

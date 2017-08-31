package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wastebot/dbs"
)

type Result struct {
	CurMonth []dbs.Result
	YearStat []dbs.OneYearStat
}

func (m *manager) genGetStat() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urlVariables := mux.Vars(r)
		uuid := urlVariables["key"]
		log.Println(uuid)
		// Ищем пользователя по uuid
		user, err := m.dbs.FindUserByUUID(uuid)
		if err != nil {
			// TODO: отправить специальный темплейт (в идеале 2: что то пошло не так и 404)
			log.Println(err)
			w.Write([]byte("Not found"))
			return
		}

		results := m.dbs.OneMonthStatistic(user.ID)
		results2 := m.dbs.OneYearStatistic(user.ID)

		result := Result{
			CurMonth: results,
			YearStat: results2,
		}

		m.pageTemplStat.Execute(w, result)
	})
}

// genFavicon возвращает фавикон.
func (m *manager) genFavicon() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: починить фавикон.
		w.Write(m.ico)
	})
}

// genPage500 генерирует страницу с сообщением о внутренней ошибке сервера.
func (m *manager) genPage500() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.page500.Execute(w, nil)
	})
}

package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wastebot/dbs"
)

// TODO: сделать i18n
var months = [12]string{"январь", "февраль", "март", "апрель", "май", "июнь", "июль", "август", "сентябрь", "октябрь", "ноябрь", "декабрь"}

type TotalStat struct {
	YearStat  map[int]*OneMonthStat
	QuickStat dbs.QuickStatistic
}

type OneMonthStat struct {
	Amount     int
	MonthTitle string
	Categories []Category
}

type Category struct {
	Category     string
	Amount       int
	Transactions *[]TransactionView1
}

type TransactionView1 struct {
	Amount    int    `json:"f1"`
	Title     string `json:"f2"`
	CreatedAt int64  `json:"f3"`
}

func (m *manager) genGetStat() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		yearStatByCategories := m.dbs.OneYearStatistic(user.ID)
		yearStat := map[int]*OneMonthStat{
			1:  &OneMonthStat{},
			2:  &OneMonthStat{},
			4:  &OneMonthStat{},
			5:  &OneMonthStat{},
			3:  &OneMonthStat{},
			6:  &OneMonthStat{},
			7:  &OneMonthStat{},
			8:  &OneMonthStat{},
			9:  &OneMonthStat{},
			10: &OneMonthStat{},
			11: &OneMonthStat{},
			12: &OneMonthStat{},
		}

		for _, catForOneMonth := range yearStatByCategories {
			log.Println("Категории по месяцам: ", catForOneMonth)
			cat := Category{
				Category: catForOneMonth.Category,
				Amount:   catForOneMonth.Amount,
			}

			transactions := &[]TransactionView1{}
			err := json.Unmarshal([]byte(catForOneMonth.RawTransactionJSON), transactions)
			if err != nil {
				log.Println(err)
			}
			cat.Transactions = transactions
			yearStat[catForOneMonth.Month].Amount += catForOneMonth.Amount
			yearStat[catForOneMonth.Month].MonthTitle = months[catForOneMonth.Month-1]
			yearStat[catForOneMonth.Month].Categories = append(yearStat[catForOneMonth.Month].Categories, cat)
		}

		quickStat := m.dbs.QuickStatistic(user.ID)

		totalStat := TotalStat{
			YearStat:  yearStat,
			QuickStat: quickStat,
		}

		log.Println(totalStat)
		m.pageTemplStat.Execute(w, totalStat)

		// Потрачено всего за день
		// Потрачено всего за месяц
		// Потрачено всего за год
		// Самая затратная категория в месяце

		// месяц - разбивка категория-потрачено денег
		// год - разбивка месяц -потрачено за месяц

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

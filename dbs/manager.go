package dbs

import (
	"encoding/json"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type MinusTransaction struct {
	Amount int    `json:"f1"`
	Title  string `json:"f2"`
}

type Result struct {
	Category string
	RawJSON  string
	Amount   int
	JSON     *[]MinusTransaction
}

type manager struct {
	db *gorm.DB
}

type Manager interface {
	OneMonthStatistic(userID int) (results []Result)
}

// NewManager - конструктор
func NewManager(db *gorm.DB) Manager {
	manager := &manager{
		db: db,
	}

	return manager
}

// OneMonthStatistic - транзакции за месяц
func (m *manager) OneMonthStatistic(userID int) (results []Result) {
	year, month, _ := time.Now().Date()
	firstDayOfTheMonthDate := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	results = []Result{}

	m.db.Raw("select category, json_agg((amount, title)) as raw_json, sum(amount) as amount from minus_transactions WHERE created_at > ? AND user_id = ? Group By category", firstDayOfTheMonthDate, userID).Scan(&results)

	log.Println(results)

	for i := range results {
		transactions := &[]MinusTransaction{}
		err := json.Unmarshal([]byte(results[i].RawJSON), transactions)
		if err != nil {
			log.Println(err)
		}
		results[i].JSON = transactions
	}
	return results

}

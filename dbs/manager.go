package dbs

import (
	"encoding/json"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Transaction struct {
	ID        uint64
	UserID    uint64
	Amount    uint64
	Title     string
	Category  string
	CreatedAt int64
	UpdatedAt int64
}

type User struct {
	ID           uint64
	TelegramID   uint64
	FirstName    string
	LastName     string
	LanguageCode string
	CreatedAt    int64
	UpdatedAt    int64
}

type MTransaction struct {
	Amount int    `json:"f1"`
	Title  string `json:"f2"`
}

type Result struct {
	Category string
	RawJSON  string
	Amount   int
	JSON     *[]MTransaction
}

type manager struct {
	db *gorm.DB
}

type Manager interface {
	OneMonthStatistic(userID int) (results []Result)
	FindOrCreateUser(tid uint64, fname string, lname string, lang string) (*User, error)
	CreateTransaction(uid uint64, amount uint64, title string, category string) (*Transaction, error)
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

	m.db.Raw("select category, json_agg((amount, title)) as raw_json, sum(amount) as amount from transactions WHERE created_at > ? AND user_id = ? Group By category", firstDayOfTheMonthDate, userID).Scan(&results)

	log.Println(results)

	for i := range results {
		transactions := &[]MTransaction{}
		err := json.Unmarshal([]byte(results[i].RawJSON), transactions)
		if err != nil {
			log.Println(err)
		}
		results[i].JSON = transactions
	}
	return results

}

// CreateTransaction - ...
func (m *manager) CreateTransaction(uid uint64, amount uint64, title string, category string) (*Transaction, error) {
	w := &Transaction{
		UserID:    uid,
		Amount:    amount,
		Title:     title,
		Category:  category,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	err := m.db.Save(w).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return w, nil
}

// CreateUser - ...
func (m *manager) createUser(tid uint64, fname string, lname string, lang string) (*User, error) {
	w := &User{
		TelegramID:   tid,
		FirstName:    fname,
		LastName:     lname,
		LanguageCode: lang,
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
	}
	err := m.db.Save(w).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return w, nil
}

// FindUser - ...
func (m *manager) findUser(tid uint64) (*User, error) {
	user := &User{}

	err := m.db.Where("telegram_id = ?", tid).First(user).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return user, nil
}

func (m *manager) FindOrCreateUser(tid uint64, fname string, lname string, lang string) (*User, error) {
	user, err := m.findUser(tid)
	if err == nil && user != nil {
		return user, nil
	}

	user, err = m.createUser(tid, fname, lname, lang)
	if err != nil {
		return nil, err
	}

	return user, nil
}

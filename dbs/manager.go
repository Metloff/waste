package dbs

import (
	"encoding/json"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Transaction struct {
	ID         uint64
	UserID     uint64
	TelegramID uint64
	Amount     uint64
	Title      string
	Category   string
	CreatedAt  int64
	UpdatedAt  int64
}

type User struct {
	ID           uint64
	TelegramID   uint64
	FirstName    string
	LastName     string
	LanguageCode string
	UUID         string
	CreatedAt    int64
	UpdatedAt    int64
}

type MTransaction struct {
	Amount    int    `json:"f1"`
	Title     string `json:"f2"`
	CreatedAt int64  `json:"f3"`
}

type ListResult struct {
	Category string
	Amount   int
	Count    int
}

type Result struct {
	Category string
	RawJSON  string
	Amount   int
	Count    int
	JSON     *[]MTransaction
}

type OneYearStat struct {
	Month  int
	Amount int
}

type manager struct {
	db *gorm.DB
}

type Manager interface {
	OneMonthStatistic(userID uint64) (results []Result)
	OneYearStatistic(userID uint64) (results []OneYearStat)
	FindOrCreateUser(tid uint64, fname string, lname string, lang string) (*User, error)
	CreateTransaction(uid uint64, tid uint64, amount uint64, title string, category string) (*Transaction, error)
	CurrentMonthStatistic(userID uint64) (results []ListResult)
	PreviousMonthStatistic(userID uint64) (results []ListResult)
	SetUUID(tid uint64, uuid string) (*User, error)
	FindUserByUUID(uuid string) (*User, error)
}

// NewManager - конструктор
func NewManager(db *gorm.DB) Manager {
	manager := &manager{
		db: db,
	}

	return manager
}

// OneMonthStatistic - транзакции за месяц
func (m *manager) OneMonthStatistic(userID uint64) (results []Result) {
	results = []Result{}

	m.db.Raw(`SELECT category, json_agg((amount, title, created_at)) as raw_json, sum(amount) as amount, count(amount) as count 
			FROM transactions 
			WHERE date_part('month', TO_TIMESTAMP(created_at)) = date_part('month', NOW()) 
			AND date_part('year', TO_TIMESTAMP(created_at)) = date_part('year', NOW())
			AND user_id = ? Group By category`, userID).Scan(&results)

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

// OneYearStatistic - количество потраченных денег по месяцам
func (m *manager) OneYearStatistic(userID uint64) (results []OneYearStat) {
	results = []OneYearStat{}

	m.db.Raw(`SELECT  date_part('month', TO_TIMESTAMP(created_at)) as month, sum(amount) as amount 
			FROM transactions 
			WHERE date_part('year', TO_TIMESTAMP(created_at)) = date_part('year', NOW())
			AND user_id = ? 
			Group By date_part('month', TO_TIMESTAMP(created_at))`, userID).Scan(&results)
	log.Println(results)

	return results
}

// CurrentMonthStatistic - транзакции за месяц
func (m *manager) CurrentMonthStatistic(userID uint64) (results []ListResult) {
	results = []ListResult{}

	m.db.Raw(`SELECT category, sum(amount) as amount, count(amount) as count
			FROM transactions 
			WHERE date_part('month', TO_TIMESTAMP(created_at)) = date_part('month', NOW())
			AND date_part('year', TO_TIMESTAMP(created_at)) = date_part('year', NOW()) 
			AND telegram_id = ? Group By category`, userID).Scan(&results)

	log.Println(results)

	return results
}

// PreviousMonthStatistic - транзакции за месяц
func (m *manager) PreviousMonthStatistic(userID uint64) (results []ListResult) {
	results = []ListResult{}

	m.db.Raw(`SELECT category, sum(amount) as amount, count(amount) as count
			FROM transactions 
			WHERE date_part('month', TO_TIMESTAMP(created_at)) = date_part('month', date_trunc('day', NOW() - interval '1 month')) 
			AND date_part('year', TO_TIMESTAMP(created_at)) = date_part('year', date_trunc('day', NOW() - interval '1 month'))
			AND telegram_id = ? Group By category`, userID).Scan(&results)

	log.Println(results)

	return results
}

// CreateTransaction - ...
func (m *manager) CreateTransaction(uid uint64, tid uint64, amount uint64, title string, category string) (*Transaction, error) {
	w := &Transaction{
		UserID:     uid,
		TelegramID: tid,
		Amount:     amount,
		Title:      title,
		Category:   category,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
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

func (m *manager) FindUserByUUID(uuid string) (*User, error) {
	user := &User{}

	err := m.db.Where("uuid = ?", uuid).First(user).Error
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

func (m *manager) SetUUID(tid uint64, uuid string) (*User, error) {
	user, err := m.findUser(tid)
	if err != nil {
		return nil, err
	}

	user.UUID = uuid
	user.UpdatedAt = time.Now().Unix()

	err = m.db.Save(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

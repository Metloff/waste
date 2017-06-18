package main

import (
	"log"

	"github.com/wastebot/api"
	"github.com/wastebot/dbs"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// type Transaction struct {}

// User - ...
type User struct {
	ID         int
	TelegramID int
	FirstName  string
	LastName   string
}

func main() {
	db := prepareDb()
	defer db.Close()

	dbsManager := dbs.NewManager(db)
	apiManger := api.NewManager(dbsManager)

	apiManger.Listen(":3000")
}

// PrepareDb - prepare postgres connection
func prepareDb() (db *gorm.DB) {
	db, err := gorm.Open("postgres", "host=localhost user=forapp dbname=trans sslmode=disable")
	if err != nil {
		log.Println(err)
	}
	return db
}

// TODO:
// Передавать флагами настройки постгреса

package main

import (
	"flag"
	"log"
	"os"

	"github.com/wastebot/app"
	"github.com/wastebot/dbs"
	"github.com/wastebot/tg"

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
	tgToken := flag.String("token", "", "Telegram bot token")
	flag.Parse()

	db := prepareDb()
	defer db.Close()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(os.Stdout)
	dbsManager := dbs.NewManager(db)
	// apiManger := api.NewManager(dbsManager)
	tgManager := tg.NewManager(*tgToken)
	appManager := app.NewManager(dbsManager, tgManager)
	appManager.Listen()
	// apiManger.Listen(":3000")
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

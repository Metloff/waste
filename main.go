package main

import (
	"fmt"
	"log"
	"os"

	bugsnag "github.com/bugsnag/bugsnag-go"
	cli "github.com/jawher/mow.cli"
	"github.com/wastebot/api"
	"github.com/wastebot/app"

	"github.com/wastebot/dbs"
	"github.com/wastebot/tg"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	application = cli.App("Waste bot", "Waste bot")

	// Telegram flags.
	tgToken = application.StringOpt("token", "", "Telegram bot token")

	// Bagsnag flags.
	bugsnagKey = application.StringOpt("bugsnug-key", "", "Specify bugsnag API key")
	bugsnagRS  = application.StringOpt("bugsnug-release-stage", "develop", "Specify bugsnag release stage")

	// PostgreSQL flags.
	pgHost    = application.StringOpt("pg-host", "127.0.0.1", "PostgreSQL host")
	pgPort    = application.IntOpt("pg-port", 5432, "PostgreSQL port")
	pgUser    = application.StringOpt("pg-user", "ma", "PostgreSQL user")
	pgDBName  = application.StringOpt("pg-db-name", "madb", "PostgreSQL database name")
	pgDBPswd  = application.StringOpt("pg-db-password", "", "PostgreSQL database password")
	pgVerbose = application.BoolOpt("pg-verbose", true, "Log PostgreSQL transactions")

	// Waste bot flags.
	apiAddr = application.StringOpt("api-addr", "127.0.0.1:3000",
		"Address of api")
	appURL = application.StringOpt("app-url", "http://127.0.0.1:3000",
		"Url for statistics")
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(os.Stdout)

	application.Action = appAction
}

func main() {
	if err := application.Run(os.Args); err != nil {
		log.Fatalln("cli:", err)
	}
}

func appAction() {
	defer bugsnag.Recover()

	bugsnag.Configure(bugsnag.Configuration{
		APIKey:       *bugsnagKey,
		ReleaseStage: *bugsnagRS,

		// To prevent process fork.
		PanicHandler: func() {},
	})

	psqlClient := preparePostgresDb(*pgHost, *pgPort, *pgUser, *pgDBName, *pgDBPswd, *pgVerbose)
	defer psqlClient.Close()

	dbsManager := dbs.NewManager(psqlClient)
	apiManger := api.NewManager(dbsManager)
	tgManager := tg.NewManager(*tgToken)
	appManager := app.NewManager(dbsManager, tgManager, *appURL)

	go func() {
		panic(apiManger.Listen(*apiAddr))
	}()
	appManager.Listen()

}

// preparePostgresDb - Подключение к PostgreSQL.
func preparePostgresDb(pgHost string, pgPort int, pgUser string, pgDBName string, pgDBPswd string, pgVerbose bool) (db *gorm.DB) {
	psqlOpener := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		pgHost, pgPort, pgUser, pgDBName)
	if pgDBPswd != "" {
		psqlOpener = fmt.Sprintf("%s password=%s", psqlOpener, pgDBPswd)
	}
	db, err := gorm.Open("postgres", psqlOpener)
	if err != nil {
		bugsnag.Notify(err, bugsnag.MetaData{
			"Error": {
				"Description": "Не могу подключиться к PostgreSQL (GORM)",
			}})
	}
	db.LogMode(pgVerbose)

	return db
}

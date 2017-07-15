package api

import (
	"log"
	"net/http"

	"github.com/wastebot/dbs"
)

type manager struct {
	dbs dbs.Manager
}

type Manager interface {
	Listen(port string) error
}

// NewManager - конструктор
func NewManager(dbs dbs.Manager) Manager {
	manager := &manager{
		dbs: dbs,
	}

	fpath := "/Users/forapp/go/src/github.com/wastebot/assets"
	http.HandleFunc("/", manager.generateHandlerGetStat())

	// GET: Ассеты
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(fpath))))

	return manager
}

// Listen - ...
func (m *manager) Listen(port string) error {
	log.Println("Listen port", port)
	return http.ListenAndServe(port, nil)
}

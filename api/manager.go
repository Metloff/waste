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
	Listen(port string)
}

// NewManager - конструктор
func NewManager(dbs dbs.Manager) Manager {
	manager := &manager{
		dbs: dbs,
	}

	fpath := "/Users/forapp/go/src/wastebot/assets"
	http.HandleFunc("/", manager.generateHandlerGetStat())

	// GET: Ассеты
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(fpath))))

	return manager
}

// Listen - ...
func (m *manager) Listen(port string) {
	log.Println("Listen port", port)
	http.ListenAndServe(port, nil)
}

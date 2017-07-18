package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wastebot/dbs"
)

type manager struct {
	r   *mux.Router
	dbs dbs.Manager
}

type Manager interface {
	Listen(port string) error
}

// NewManager - конструктор
func NewManager(dbs dbs.Manager) Manager {
	r := mux.NewRouter()

	manager := &manager{
		r:   r,
		dbs: dbs,
	}

	fpath := "/Users/forapp/go/src/github.com/wastebot/assets"
	// GET: Ассеты
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir(fpath))))

	r.HandleFunc("/{key}", manager.generateHandlerGetStat()).Methods("GET")

	// http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(fpath))))

	return manager
}

// Listen - функция запуска сервера.
func (m *manager) Listen(port string) error {
	srv := &http.Server{
		Handler: m.r,
		Addr:    port,
	}

	log.Println("Server starting on", port)

	return srv.ListenAndServe()
}

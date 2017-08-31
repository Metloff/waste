package api

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wastebot/dbs"
)

type manager struct {
	r             *mux.Router
	dbs           dbs.Manager
	pageTemplStat *template.Template
}

type Manager interface {
	Listen(port string) error
}

// NewManager - конструктор
func NewManager(dbs dbs.Manager, htmlPageStat []byte) Manager {
	r := mux.NewRouter()

	manager := &manager{
		r:   r,
		dbs: dbs,
	}

	manager.pageTemplStat = template.Must(template.New("page_stat.html").
		Parse(string(htmlPageStat)))

	fpath := "/Users/forapp/go/src/github.com/wastebot/assets"
	// GET: Ассеты
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir(fpath))))

	r.HandleFunc("/{key}", manager.generateHandlerGetStat()).Methods("GET")

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

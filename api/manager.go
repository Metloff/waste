package api

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"runtime/debug"

	bugsnag "github.com/bugsnag/bugsnag-go"
	"github.com/gorilla/mux"
	"github.com/wastebot/dbs"
)

type manager struct {
	router        *mux.Router
	dbs           dbs.Manager
	pageTemplStat *template.Template
	page500       *template.Template
	ico           []byte
}

type Manager interface {
	Listen(port string) error
}

// NewManager - конструктор
func NewManager(dbs dbs.Manager, htmlPageStat, htmlPage500, ico []byte) Manager {
	m := &manager{
		router: mux.NewRouter(),
		dbs:    dbs,
		ico:    ico,
	}

	// Templates.
	m.pageTemplStat = template.Must(template.New("page_stat.html").
		Parse(string(htmlPageStat)))

	m.page500 = template.Must(template.New("page_500.html").
		Parse(string(htmlPage500)))

	// Routes.
	m.router.
		Methods("GET").
		Path("/favicon.ico").
		Name("Favicon").
		Handler(m.genFavicon())

	m.router.
		Methods("GET").
		Path("/page500").
		Name("Favicon").
		Handler(m.genPage500())

	m.router.
		Methods("GET").
		Path("/{key}").
		Name("GetStat").
		Handler(m.wrapHTMLRec(m.genGetStat()))

	return m
}

// Listen - функция запуска сервера.
func (m *manager) Listen(port string) error {
	srv := &http.Server{
		Handler: m.router,
		Addr:    port,
	}

	log.Println("Server starting on", port)

	return srv.ListenAndServe()
}

// wrapHTMLRec в случае возникновения необработанной паники, обрабатывает панику
// и перенаправляе пользователя на страницу 500.
func (m *manager) wrapHTMLRec(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			rec := recover()
			if rec != nil {
				switch t := rec.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("panic with unknown error")
				}

				bugsnag.Notify(err, map[string]interface{}{
					"stacktrace": string(debug.Stack()),
					"uri":        r.RequestURI,
				})

				http.Redirect(w, r, "/page500", 301)
			}
		}()
		h.ServeHTTP(w, r)
	})
}

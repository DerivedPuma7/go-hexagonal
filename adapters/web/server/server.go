package server

import (
	"github.com/DerivedPuma7/go-hexagonal/adapters/web/handler"
	"github.com/DerivedPuma7/go-hexagonal/application/interfaces"

	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type Webserver struct {
	Service interfaces.ProductServiceInterface
}

func MakeNewWebserver() *Webserver {
	return &Webserver{}
}

func (w Webserver) Serve() {
	r := mux.NewRouter() // routes -> similar to express
	n := negroni.New( // middlewares
		negroni.NewLogger(),
	)

	handler.MakeProductHandlers(r, n, w.Service)
	http.Handle("/", r)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr: ":8080",
		Handler: http.DefaultServeMux,
		ErrorLog: log.New(os.Stderr, "log: ", log.Lshortfile),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

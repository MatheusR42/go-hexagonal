package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/matheusr42/go-hexagonal/application"
	"github.com/matheusr42/go-hexagonal/web/server/handler"
	"github.com/urfave/negroni"
)

type Webserver struct {
	Service application.ProductServiceInterface
}

func MakeNewWebserver(service application.ProductServiceInterface) *Webserver {
	return &Webserver{}
}

func (w Webserver) Serve() {
	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandlers(r, n, w.Service)
	http.Handle("/", r)

	server := &http.Server{
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		Addr:              ":8080",
		Handler:           nil,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

package server

import (
	"github.com/mrdibre/hexagonal-arch-go/adapters/web/handler"
	"github.com/mrdibre/hexagonal-arch-go/application"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type WebServer struct {
	 Service application.ProductServiceInterface
}

func MakeNewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {
	router := mux.NewRouter()
	n := negroni.New(negroni.NewLogger())

	handler.MakeProductHandler(router, n, w.Service)
	http.Handle("/", router)

	server := &http.Server{
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr: ":9000",
		Handler: http.DefaultServeMux,
		ErrorLog: log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
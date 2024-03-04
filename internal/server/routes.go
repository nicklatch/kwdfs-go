package server

import (
	"log"
	"net/http"
	"nicklatch/kwdfs-go/cmd/web"
)

type PageData struct {
	Title    string
	Endpoint string
}

func (s *Server) RegisteredRoutes() http.Handler {

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.FS(web.Files))
	mux.Handle("GET /assets/*", fileServer)

	mux.HandleFunc("GET /login", s.loginHandler)

	mux.HandleFunc("GET /", s.rootHandler)
	mux.HandleFunc("GET /api/index", s.indexRootContentHandler)

	mux.HandleFunc("GET /dealers", s.dealerHandler)
	mux.HandleFunc("GET /api/dealers", s.dealersGetTable)

	mux.HandleFunc("GET /locations", s.locationHandler)
	mux.HandleFunc("GET /api/locations", s.locationsGetTable)

	mux.HandleFunc("GET /customers", s.customerHandler)
	mux.HandleFunc("GET /api/customers", s.customersGetTable)
	mux.HandleFunc("GET /api/customers/{name}", s.customersGetOne)

	return mux
}

func (s *Server) loginHandler(rw http.ResponseWriter, _ *http.Request) {
	data := struct {
		Title string
	}{Title: "Login"}
	err := s.tmpl.ExecuteTemplate(rw, "login.html", data)
	if err != nil {
		log.Print(err)
	}
}

func (s *Server) rootHandler(rw http.ResponseWriter, _ *http.Request) {
	data := PageData{
		Title:    "Home",
		Endpoint: "index",
	}
	err := s.tmpl.ExecuteTemplate(rw, "base.html", data)
	if err != nil {
		log.Print(err)
	}
}

func (s *Server) indexRootContentHandler(rw http.ResponseWriter, _ *http.Request) {
	err := s.tmpl.ExecuteTemplate(rw, "index-content", nil)
	if err != nil {
		log.Print(err)
	}
}

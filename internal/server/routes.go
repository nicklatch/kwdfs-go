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
	mux.Handle("/assets/*", fileServer)

	mux.HandleFunc("/login", s.loginHandler)

	mux.HandleFunc("/", s.rootHandler)
	mux.HandleFunc("/api/index", s.indexRootContentHandler)

	mux.HandleFunc("/dealers", s.dealerHandler)
	mux.HandleFunc("/api/dealers", s.dealersGetTable)

	mux.HandleFunc("/locations", s.locationHandler)
	mux.HandleFunc("/api/locations", s.locationsGetTable)

	mux.HandleFunc("/customers", s.customerHandler)
	mux.HandleFunc("/api/customers", s.customersGetTable)

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

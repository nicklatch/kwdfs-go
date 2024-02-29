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

	mux.HandleFunc("/login", s.LoginHandler)

	mux.HandleFunc("/", s.RootHandler)
	mux.HandleFunc("/api/index", s.IndexRootContentHandler)

	mux.HandleFunc("/dealers", s.DealerHandler)
	mux.HandleFunc("/api/dealers", s.DealersGetTable)

	mux.HandleFunc("/locations", s.LocationHandler)
	mux.HandleFunc("/api/locations", s.LocationsGetTable)

	mux.HandleFunc("/customers", s.CustomerHandler)
	mux.HandleFunc("/api/customers", s.CustomersGetTable)

	return mux
}

func (s *Server) LoginHandler(rw http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
	}{Title: "Login"}
	err := s.tmpl.ExecuteTemplate(rw, "login.html", data)
	if err != nil {
		log.Print(err)
	}
}

func (s *Server) RootHandler(rw http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:    "Home",
		Endpoint: "index",
	}
	err := s.tmpl.ExecuteTemplate(rw, "index.html", data)
	if err != nil {
		log.Print(err)
	}
}

func (s *Server) IndexRootContentHandler(rw http.ResponseWriter, r *http.Request) {
	err := s.tmpl.ExecuteTemplate(rw, "index-content.html", nil)
	if err != nil {
		log.Print(err)
	}
}

func (s *Server) CustomerHandler(rw http.ResponseWriter, r *http.Request) {
	data := PageData{Title: "Customers", Endpoint: "customers"}
	err := s.tmpl.ExecuteTemplate(rw, "index.html", data)
	if err != nil {
		log.Print(err)
	}
}

func (s *Server) CustomersGetTable(rw http.ResponseWriter, r *http.Request) {
	err := s.tmpl.ExecuteTemplate(rw, "customers-table.html", "Customers")
	if err != nil {
		log.Print(err)
	}
}

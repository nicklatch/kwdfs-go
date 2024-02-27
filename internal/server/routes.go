package server

import (
	"log"
	"net/http"
	"nicklatch/kwdfs-go/cmd/web"
)

func (s *Server) RegisteredRoutes() http.Handler {

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.FS(web.Files))
	mux.Handle("/assets/*", fileServer)
	mux.HandleFunc("/", s.RootHandler)
	mux.HandleFunc("/login", s.LoginHandler)

	mux.HandleFunc("/dealers", s.DealerHandler)
	mux.HandleFunc("/api/dealers", s.DealerGetTable)

	mux.HandleFunc("/locations", s.LocationHandler)
	mux.HandleFunc("/api/locations", s.LocationsGetTable)
	mux.HandleFunc("/customers", s.CustomerHandler)
	mux.HandleFunc("/api/customers", s.CustomersGetTable)

	return mux
}

func (s *Server) RootHandler(rw http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
	}{Title: "Home"}
	err := s.tmpl.ExecuteTemplate(rw, "index.html", data)
	if err != nil {
		log.Print(err)
	}
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

func (s *Server) DealerHandler(rw http.ResponseWriter, r *http.Request) {
	data := struct {
		Title    string
		Endpoint string
	}{Title: "Dealers", Endpoint: "dealers"}
	err := s.tmpl.ExecuteTemplate(rw, "index.html", data)
	if err != nil {
		log.Print(err)
	}
}

func (s *Server) DealerGetTable(rw http.ResponseWriter, r *http.Request) {
	err := s.tmpl.ExecuteTemplate(rw, "dealer-content.html", "")
	if err != nil {
		log.Print(err)
	}
}

func (s *Server) LocationHandler(rw http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
	}{Title: "Locations"}
	err := s.tmpl.ExecuteTemplate(rw, "index.html", data)
	if err != nil {
		log.Print(err)
	}
}

func (s *Server) LocationsGetTable(rw http.ResponseWriter, r *http.Request) {
	err := s.tmpl.ExecuteTemplate(rw, "locations-content.html", "")
	if err != nil {
		log.Print(err)
	}
}

func (s *Server) CustomerHandler(rw http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
	}{Title: "Customers"}
	err := s.tmpl.ExecuteTemplate(rw, "index.html", data)
	if err != nil {
		log.Print(err)
	}
}

func (s *Server) CustomersGetTable(rw http.ResponseWriter, r *http.Request) {
	err := s.tmpl.ExecuteTemplate(rw, "customers-content.html", "")
	if err != nil {
		log.Print(err)
	}
}

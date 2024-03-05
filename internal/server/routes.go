package server

import (
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

	return mux
}

func (s *Server) loginHandler(rw http.ResponseWriter, req *http.Request) {
	data := struct {
		Title string
	}{Title: "Login"}
	err := s.tmpl.ExecuteTemplate(rw, "login.gohtml", data)
	if err != nil {
		s.log.Error(err)
	}

	s.log.Infof("%v", req.Method)
}

func (s *Server) rootHandler(rw http.ResponseWriter, req *http.Request) {
	data := PageData{
		Title:    "Home",
		Endpoint: "index",
	}
	err := s.tmpl.ExecuteTemplate(rw, "base.gohtml", data)
	if err != nil {
		s.log.Error(err)
	}
	s.log.Infof("%v", req.Method)
}

func (s *Server) indexRootContentHandler(rw http.ResponseWriter, req *http.Request) {
	err := s.tmpl.ExecuteTemplate(rw, "index-content", nil)
	if err != nil {
		s.log.Error(err)
	}

	s.log.Infof("%v", req.Method)
}

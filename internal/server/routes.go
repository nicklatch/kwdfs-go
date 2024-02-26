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
	mux.HandleFunc("/dealers", s.RootHandler)
	mux.HandleFunc("/locations", s.RootHandler)
	mux.HandleFunc("/customers", s.RootHandler)

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

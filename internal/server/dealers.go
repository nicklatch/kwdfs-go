package server

import (
	"log"
	"net/http"
	"nicklatch/kwdfs-go/internal/db"
)

func (s *Server) DealerHandler(rw http.ResponseWriter, r *http.Request) {
	data := PageData{Title: "Dealers", Endpoint: "dealers"}
	err := s.tmpl.ExecuteTemplate(rw, "index.html", data)
	if err != nil {
		log.Print(err)
	}
}

func (s *Server) DealersGetTable(rw http.ResponseWriter, r *http.Request) {
	err := s.tmpl.ExecuteTemplate(rw, "dealers-table", db.GetAllDealerGroups())
	if err != nil {
		log.Print(err)
	}
}

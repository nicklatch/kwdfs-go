package server

import (
	"context"
	"log"
	"net/http"
)

func (s *Server) dealerHandler(rw http.ResponseWriter, r *http.Request) {
	data := PageData{Title: "Dealers", Endpoint: "dealers"}
	err := s.tmpl.ExecuteTemplate(rw, "base.html", data)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(500)
	}
	log.Println(r)
}

func (s *Server) dealersGetTable(rw http.ResponseWriter, r *http.Request) {
	data, err := s.db.GetAllDealers(context.Background())
	if err != nil {
		log.Println(err)
		rw.WriteHeader(400) // TODO: tmpl fragment to return
	}

	err = s.tmpl.ExecuteTemplate(rw, "dealers-table", data)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(500) // TODO: tmpl fragment to return
	}
	log.Printf("dealersGetTable: Method: %v", r.Method)
}

package server

import (
	"context"
	"log"
	"net/http"
)

func (s *Server) customerHandler(rw http.ResponseWriter, r *http.Request) {
	data := PageData{Title: "Customers", Endpoint: "customers"}
	err := s.tmpl.ExecuteTemplate(rw, "base.html", data)
	if err != nil {
		log.Print(err)
	}
}

func (s *Server) customersGetTable(rw http.ResponseWriter, _ *http.Request) {
	cd, err := s.db.GetAllCustomers(context.Background())
	if err != nil {
		log.Println(err)
		rw.WriteHeader(400) // TODO: tmpl fragment to return
	}

	err = s.tmpl.ExecuteTemplate(rw, "customers-table", cd)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(500) // TODO: tmpl fragment to return
	}
}

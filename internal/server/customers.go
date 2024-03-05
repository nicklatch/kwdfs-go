package server

import (
	"context"
	"github.com/charmbracelet/log"
	"net/http"
)

func (s *Server) customerHandler(rw http.ResponseWriter, req *http.Request) {
	data := PageData{Title: "Customers", Endpoint: "customers"}
	err := s.tmpl.ExecuteTemplate(rw, "base.gohtml", data)
	if err != nil {
		log.Print(err)
	}

	s.log.Infof("%v", req.Method)
}

func (s *Server) customersGetTable(rw http.ResponseWriter, req *http.Request) {
	cd, err := s.db.GetAllCustomers(context.Background())
	if err != nil {
		s.log.Error(err)
		rw.WriteHeader(400) // TODO: tmpl fragment to return
	}

	err = s.tmpl.ExecuteTemplate(rw, "customers-table", cd)
	if err != nil {
		s.log.Error(err)
		rw.WriteHeader(500) // TODO: tmpl fragment to return
	}

	s.log.Infof("%v", req.Method)
}

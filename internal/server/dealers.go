package server

import (
	"context"
	"net/http"
)

func (s *Server) dealerHandler(rw http.ResponseWriter, req *http.Request) {
	data := PageData{Title: "Dealers", Endpoint: "dealers"}
	err := s.tmpl.ExecuteTemplate(rw, "base.gohtml", data)
	if err != nil {
		s.log.Error(err)
		rw.WriteHeader(500)
	}
	s.log.Infof("'/dealers' : %v", req.Method)
}

func (s *Server) dealersGetTable(rw http.ResponseWriter, req *http.Request) {
	data, err := s.db.GetAllDealers(context.Background())
	if err != nil {
		s.log.Error(err)
		rw.WriteHeader(400) // TODO: tmpl fragment to return
	}

	err = s.tmpl.ExecuteTemplate(rw, "dealers-table", data)
	if err != nil {
		s.log.Error(err)
		rw.WriteHeader(500) // TODO: tmpl fragment to return
	}
	s.log.Infof("%v", req.Method)
}

package server

import (
	"context"
	"net/http"
)

func (s *Server) locationHandler(rw http.ResponseWriter, req *http.Request) {
	data := PageData{
		Title:    "Locations",
		Endpoint: "locations",
	}
	err := s.tmpl.ExecuteTemplate(rw, "base.gohtml", data)
	if err != nil {
		s.log.Error(err)
	}

	s.log.Infof("%v", req.Method)
}

func (s *Server) locationsGetTable(rw http.ResponseWriter, req *http.Request) {
	data, err := s.db.GetAllLocations(context.Background())
	if err != nil {
		s.log.Error(err)
		rw.WriteHeader(400) // TODO: tmpl fragment to return
	}

	err = s.tmpl.ExecuteTemplate(rw, "locations-table", data)
	if err != nil {
		s.log.Error(err)
		rw.WriteHeader(500) // TODO: tmpl fragment to return
	}
	s.log.Infof("%v", req.Method)
}

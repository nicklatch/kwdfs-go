package server

import (
	"context"
	"log"
	"net/http"
)

func (s *Server) locationHandler(rw http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:    "Locations",
		Endpoint: "locations",
	}
	err := s.tmpl.ExecuteTemplate(rw, "base.html", data)
	if err != nil {
		log.Fatalln(err)
	}
}

func (s *Server) locationsGetTable(rw http.ResponseWriter, r *http.Request) {
	data, err := s.db.GetAllLocations(context.Background())
	if err != nil {
		log.Println(err)
		rw.WriteHeader(400) // TODO: tmpl fragment to return
	}

	err = s.tmpl.ExecuteTemplate(rw, "locations-table", data)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(500) // TODO: tmpl fragment to return
	}

	log.Printf("locationsGetTable: Method=%v", r.Method)
}

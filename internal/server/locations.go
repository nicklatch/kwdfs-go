package server

import (
	"log"
	"net/http"
	"nicklatch/kwdfs-go/internal/db"
)

func (s *Server) LocationHandler(rw http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:    "Locations",
		Endpoint: "locations",
	}
	err := s.tmpl.ExecuteTemplate(rw, "index.html", data)
	if err != nil {
		log.Print(err)
	}
}

// This will need to utilize a join to get the dealer name from
// their respective UUID

func (s *Server) LocationsGetTable(rw http.ResponseWriter, r *http.Request) {
	err := s.tmpl.ExecuteTemplate(rw, "locations-table", db.GetAllLocations())
	if err != nil {
		log.Print(err)
	}
}

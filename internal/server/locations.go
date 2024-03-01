package server

import (
	"log"
	"net/http"
)

func (s *Server) LocationHandler(rw http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:    "Locations",
		Endpoint: "locations",
	}
	err := s.tmpl.ExecuteTemplate(rw, "index.html", data)
	if err != nil {
		log.Fatalln(err)
	}
}

func (s *Server) LocationsGetTable(rw http.ResponseWriter, r *http.Request) {
	data, err := s.db.GetAllLocations()
	if err != nil {
		log.Println(err)
		rw.WriteHeader(400) // TODO: tmpl fragment to return
	}

	err = s.tmpl.ExecuteTemplate(rw, "locations-table", data)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(500) // TODO: tmpl fragment to return
	}

	log.Printf("LocationsGetTable: Method=%v", r.Method)
}

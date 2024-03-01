package server

import (
	"log"
	"net/http"
)

func (s *Server) CustomerHandler(rw http.ResponseWriter, r *http.Request) {
	data := PageData{Title: "Customers", Endpoint: "customers"}
	err := s.tmpl.ExecuteTemplate(rw, "index.html", data)
	if err != nil {
		log.Print(err)
	}
}

func (s *Server) CustomersGetTable(rw http.ResponseWriter, r *http.Request) {
	data, err := s.db.GetAllCustomers()
	if err != nil {
		log.Println(err)
		rw.WriteHeader(400) // TODO: tmpl fragment to return
	}
	err = s.tmpl.ExecuteTemplate(rw, "customers-table", data)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(500) // TODO: tmpl fragment to return
	}
}

package server

import (
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
	cd, err := s.db.GetAllCustomers()
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

func (s *Server) customersGetOne(rw http.ResponseWriter, r *http.Request) {
	cust, err := s.db.GetOneCustomer(r.PathValue("name"))

	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
	}

	log.Println(cust)

	// returns no-content(204) because htmx will swap in the information on the frontend
	err = s.tmpl.ExecuteTemplate(rw, "customers-table", cust)
}

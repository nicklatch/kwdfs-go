package server

import (
	"log"
	"net/http"
	"nicklatch/kwdfs-go/internal/database"
	"time"
)

var templCache = make(map[time.Time][]database.Dealer)

func (s *Server) DealerHandler(rw http.ResponseWriter, r *http.Request) {
	data := PageData{Title: "Dealers", Endpoint: "dealers"}
	err := s.tmpl.ExecuteTemplate(rw, "index.html", data)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(500)
	}
	log.Println(r)
}

func (s *Server) DealersGetTable(rw http.ResponseWriter, r *http.Request) {
	data, err := s.db.GetAllDealerGroups()
	if err != nil {
		log.Println(err)
		rw.WriteHeader(400) // TODO: tmpl fragment to return
	}

	err = s.tmpl.ExecuteTemplate(rw, "dealers-table", data) // TODO: need to add db service back in here
	if err != nil {
		log.Println(err)
		rw.WriteHeader(500) // TODO: tmpl fragment to return
	}
	log.Printf("DealersGetTable: Method: %v", r.Method)
}

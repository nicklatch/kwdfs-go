package server

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"html/template"
	"net/http"
	db2 "nicklatch/kwdfs-go/internal/db"
	"os"
	"strconv"
	"time"
)

var files = []string{
	// pages
	"cmd/web/templates/pages/index.html",
	"cmd/web/templates/pages/login.html",
	// shared components
	"cmd/web/templates/components/shared/base-head.html",
	"cmd/web/templates/components/shared/header.html",
	"cmd/web/templates/components/shared/footer.html",
	"cmd/web/templates/components/shared/content.html",
	// tables and other content components
	"cmd/web/templates/components/index-content.html",
	"cmd/web/templates/components/dealers-table.html",
	"cmd/web/templates/components/locations-table.html",
	"cmd/web/templates/components/customers-table.html",
	"cmd/web/templates/components/maps.html",
}

type Server struct {
	port int
	tmpl *template.Template
	db   db2.Service
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,
		tmpl: template.Must(template.ParseFiles(files...)),
		db:   db2.New(),
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisteredRoutes(),
		IdleTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

package server

import (
	"fmt"
	"github.com/charmbracelet/log"
	_ "github.com/joho/godotenv/autoload"
	"html/template"
	"net/http"
	database "nicklatch/kwdfs-go/internal/database/sqlc"
	"os"
	"strconv"
	"time"
)

var files = []string{
	// pages
	"cmd/web/templates/pages/base.gohtml",
	"cmd/web/templates/pages/login.gohtml",

	// shared components
	"cmd/web/templates/components/shared/base-head.gohtml",
	"cmd/web/templates/components/shared/header.gohtml",
	"cmd/web/templates/components/shared/footer.gohtml",
	"cmd/web/templates/components/shared/content.gohtml",
	"cmd/web/templates/components/shared/close.gohtml",
	"cmd/web/templates/components/shared/table-buttons.gohtml",

	// tables and other content components
	"cmd/web/templates/components/index-content.gohtml",
	"cmd/web/templates/components/dealers-table.gohtml",
	"cmd/web/templates/components/locations-table.gohtml",
	"cmd/web/templates/components/customers-table.gohtml",
	"cmd/web/templates/components/maps.gohtml",
	"cmd/web/templates/components/dealer_edit.gohtml",
}

type Server struct {
	port int
	tmpl *template.Template
	db   *database.Queries
	log  *log.Logger
}

func NewServer(db *database.Queries) *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,
		tmpl: template.Must(template.ParseFiles(files...)),
		db:   db,
		log: log.NewWithOptions(os.Stderr, log.Options{
			ReportCaller:    true,
			ReportTimestamp: true,
			TimeFormat:      time.Kitchen,
		}),
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisteredRoutes(),
		IdleTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

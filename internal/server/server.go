package server

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"time"
)

var files = []string{
	"cmd/web/templates/pages/index.html",
	"cmd/web/templates/pages/login.html",
	"cmd/web/templates/components/base-head.html",
	"cmd/web/templates/components/header.html",
	"cmd/web/templates/components/footer.html",
	"cmd/web/templates/components/dealer-content.html",
	"cmd/web/templates/components/locations-content.html",
	"cmd/web/templates/components/customers-content.html",
	"cmd/web/templates/components/content.html",
}

type Server struct {
	port int
	tmpl *template.Template
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,
		tmpl: template.Must(template.ParseFiles(files...)),
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisteredRoutes(),
		IdleTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

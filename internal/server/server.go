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

var files = []string{"cmd/web/templates/index.html",
	"cmd/web/templates/base-head.html",
	"cmd/web/templates/header.html",
	"cmd/web/templates/footer.html",
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

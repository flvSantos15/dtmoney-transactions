package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router chi.Router
	Handlers map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(webServerPort string) *WebServer {
	return &WebServer{
		Router: chi.NewRouter(),
		Handlers: make(map[string]http.HandlerFunc),
		WebServerPort: webServerPort,
	}
}

func (server *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	server.Handlers[path] = handler
}

func (server *WebServer) Start() {
	server.Router.Use(middleware.Logger)
	for path, handler := range server.Handlers {
		server.Router.Post(path, handler)
	}
	http.ListenAndServe(server.WebServerPort, server.Router)
}
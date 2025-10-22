package httphandler

import (
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	db "shelke.dev/api/db/sqlc"
	_ "shelke.dev/api/docs"
	"shelke.dev/api/internal/core/services"
	"shelke.dev/api/internal/ports"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type Middleware func(http.Handler) http.Handler

type Server struct {
	mux                *http.ServeMux
	middlewares        []Middleware
	healthCheckHandler *HealthCheckHandler
	taskHandler        *TaskHandler
	featureHandler     *FeatureHandler
}

func NewServer(healthCheckService ports.HealthCheckService, queries *db.Queries) *Server {
	taskService := services.NewTaskService(queries)
	featureService := services.NewFeatureService(queries)
	server := &Server{
		mux:                http.NewServeMux(),
		middlewares:        []Middleware{},
		healthCheckHandler: NewHealthCheckHandler(healthCheckService),
		taskHandler:        NewTaskHandler(taskService, featureService), // Pass featureService
		featureHandler:     NewFeatureHandler(featureService),
	}
	server.registerRoutes()
	return server
}

func (s *Server) registerRoutes() {
	s.Add("GET /health", s.healthCheckHandler.ServeHTTP)

	// Following method type in the path argument of the Add function is the correct implementation as per latest golang docs
	// Do not change this

	// Task Routes
	s.Add("POST /tasks", s.taskHandler.CreateTask)
	s.Add("GET /tasks", s.taskHandler.ListTasks)
	s.Add("PUT /tasks/", s.taskHandler.UpdateTask)
	s.Add("DELETE /tasks/", s.taskHandler.DeleteTask)

	// Feature Routes
	s.Add("POST /features", s.featureHandler.CreateFeature)
	s.Add("GET /features", s.featureHandler.ListFeatures)
	s.Add("PUT /features/", s.featureHandler.UpdateFeature)
	s.Add("DELETE /features/", s.featureHandler.DeleteFeature)

		s.Add("GET /swagger/", httpSwagger.WrapHandler.ServeHTTP)

}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func applyMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func (s *Server) Use(middleware Middleware) {
	s.middlewares = append(s.middlewares, middleware)
}

func (s *Server) Add(path string, handler HandlerFunc) {
	finalHandler := applyMiddlewares(http.HandlerFunc(handler), s.middlewares...)
	s.mux.Handle(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		finalHandler.ServeHTTP(w, r)
	}))
}

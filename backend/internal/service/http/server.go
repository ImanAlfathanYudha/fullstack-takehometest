package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"backend/internal/middleware"
	"backend/internal/openapigen"

	"github.com/go-chi/chi/v5"
	oapinethttpmw "github.com/oapi-codegen/nethttp-middleware"
)

type Server struct {
	router http.Handler
}

const (
	readTimeout  = 10
	writeTimeout = 10
	idleTimeout  = 60
)

func NewServer(apiHandler openapigen.ServerInterface, openapiYamlPath string, jwtSecret []byte) *Server {
	swagger, err := openapigen.GetSwagger()
	if err != nil {
		log.Fatalf("failed to load swagger: %v", err)
	}
	r := chi.NewRouter()
	r.Route("/", func(api chi.Router) {
		api.Use(oapinethttpmw.OapiRequestValidator(swagger))
		openapigen.HandlerFromMux(apiHandler, api)
	})
	r.Group(func(r chi.Router) {
		r.Use(middleware.JWTAuth(jwtSecret)) // Use the jwtSecret here!
		// GET /payments
		r.Get("/dashboard/v1/payments", func(w http.ResponseWriter, req *http.Request) {
			// We manually parse query params here (status, id)
			var params openapigen.GetDashboardV1PaymentsParams
			if v := req.URL.Query().Get("status"); v != "" {
				params.Status = &v
			}
			if v := req.URL.Query().Get("id"); v != "" {
				params.Id = &v
			}
			apiHandler.GetDashboardV1Payments(w, req, params)
		})
	})
	return &Server{
		router: r,
	}
}

func (s *Server) Start(addr string) {
	service := &http.Server{
		Addr:         addr,
		Handler:      s.router,
		ReadTimeout:  readTimeout * time.Second,
		WriteTimeout: writeTimeout * time.Second,
		IdleTimeout:  idleTimeout * time.Second,
	}
	go func() {
		log.Printf("listening on %s", addr)
		err := service.ListenAndServe()
		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	log.Println("Shutting down gracefully...")

	// Timeout for shutdown
	const shutdownTimeout = 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := service.Shutdown(ctx); err != nil {
		log.Fatalf("Forced shutdown: %v", err)
	}

	log.Println("Server stopped cleanly ✔")
}

func (s *Server) Routes() http.Handler {
	return s.router
}

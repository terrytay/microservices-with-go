package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/terrytay/microservices-with-go/product-api/data"
	"github.com/terrytay/microservices-with-go/product-api/handlers"
)

func main() {
	// logger and validation packages
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	v := data.NewValidation()

	r := chi.NewRouter()

	// CORS settings
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, // Use this to allow specific origin hosts
		// AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Content-Type: application/json
	r.Use(middlewareContentJSON)

	// Route to check health status
	r.Route("/health", func(r chi.Router) {
		hh := handlers.NewHealth(l)
		r.Get("/", hh.HealthCheck)
	})

	// Route for products
	r.Route("/products", func(r chi.Router) {
		ph := handlers.NewProducts(l, v)

		r.Get("/", ph.GetProducts)                 // GET /products
		r.Get("/{id:[0-9]+}", ph.GetProduct)       // GET /product
		r.Delete("/{id:[0-9]+}", ph.DeleteProduct) // DELETE /products/:id

		r.Route("/", func(r chi.Router) {
			r.Use(ph.MiddlwareValidateProduct) // Validates body JSON format
			r.Post("/", ph.AddProduct)         // POST /products
			r.Put("/", ph.UpdateProducts)      // PUT /products
		})
	})

	// Specify file location of swagger.yaml, expose it to www
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	// Expose /docs to www, documentation for product API
	r.Handle("/docs", sh)
	r.Handle("/swagger.yaml", http.FileServer(http.Dir("./"))) // Serving file request

	// Server setup
	s := &http.Server{
		Addr:         ":9090",           // configure the bind address
		Handler:      r,                 // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		IdleTimeout:  120 * time.Second, // max time for connections using TCP keep-alive
		ReadTimeout:  1 * time.Second,   // max time to read request from the client
		WriteTimeout: 1 * time.Second,   // max time to write response to the client
	}

	// Goroutine to start up server
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// Listen to interrupt signal on channel
	sigChan := make(chan os.Signal, 10)
	signal.Notify(sigChan, os.Interrupt)

	// Shutdown gracefully upon receiving SIGINT
	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	// timeout context to pass to shutdown process
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s.Shutdown(tc)
}

func middlewareContentJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

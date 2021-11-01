package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-openapi/runtime/middleware"
	gohandlers "github.com/gorilla/handlers"
	"github.com/terrytay/microservices-with-go/product-api/data"
	"github.com/terrytay/microservices-with-go/product-api/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	v := data.NewValidation()

	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:3000"}))
	r := chi.NewRouter()

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

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	r.Handle("/docs", sh)
	r.Handle("/swagger.yaml", http.FileServer(http.Dir("./"))) // Serving file request

	s := &http.Server{
		Addr:         ":9090",           // configure the bind address
		Handler:      ch(r),             // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		IdleTimeout:  120 * time.Second, // max time for connections using TCP keep-alive
		ReadTimeout:  1 * time.Second,   // max time to read request from the client
		WriteTimeout: 1 * time.Second,   // max time to write response to the client
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 10)
	signal.Notify(sigChan, os.Interrupt)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(tc)
}

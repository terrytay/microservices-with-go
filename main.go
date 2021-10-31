package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/terrytay/product-api/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	r := chi.NewRouter()

	r.Route("/products", func(r chi.Router) {
		ph := handlers.NewProducts(l)

		r.Get("/", ph.GetProducts) // GET /products

		r.Route("/", func(r chi.Router) {
			r.Use(ph.MiddlwareValidateProduct) // Validates body JSON format
			r.Post("/", ph.AddProduct)         // POST /products
			r.Put("/{id}", ph.UpdateProducts)  // PUT /products/:id
		})
	})

	s := &http.Server{
		Addr:         ":9090",
		Handler:      r,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
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

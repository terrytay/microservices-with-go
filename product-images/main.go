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
	"github.com/joho/godotenv"
	"github.com/terrytay/microservices-with-go/product-images/files"
	"github.com/terrytay/microservices-with-go/product-images/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-images", log.LstdFlags)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	basePath := os.Getenv("BASEPATH")

	s, err := files.NewLocal(basePath, -1)
	if err != nil {
		l.Fatal("unable to create storage", err)
		os.Exit(1)
	}

	fh := handlers.NewFiles(l, s, chi.URLParam)

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

	r.Post("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}", fh.ServeHTTP)

	r.Method(http.MethodGet, "/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}", http.StripPrefix("/images/", http.FileServer(http.Dir(basePath))))

	server := &http.Server{
		Addr:         ":9090",           // configure the bind address
		Handler:      r,                 // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		IdleTimeout:  120 * time.Second, // max time for connections using TCP keep-alive
		ReadTimeout:  1 * time.Second,   // max time to read request from the client
		WriteTimeout: 1 * time.Second,   // max time to write response to the client
	}

	// Goroutine to start up server
	go func() {
		err := server.ListenAndServe()
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

	server.Shutdown(tc)
}

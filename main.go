package main

import (
	"github.com/go-chi/chi"
	"github.com/dimiro1/health"
	"net/http"
)

func main() {
	h := health.NewHandler()

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	r.Get("/health", h.ServeHTTP)

	http.ListenAndServe(":9090", r)
}

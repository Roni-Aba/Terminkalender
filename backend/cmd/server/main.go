package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Service struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	DurationMinutes int    `json:"durationMinutes"`
	Price           int    `json:"price"`
}

type Staff struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	// demo data
	services := []Service{
		{ID: 1, Name: "Haarschnitt", DurationMinutes: 30, Price: 25},
		{ID: 2, Name: "Bart trimmen", DurationMinutes: 15, Price: 12},
		{ID: 3, Name: "Färben", DurationMinutes: 90, Price: 65},
	}

	staff := []Staff{
		{ID: 1, Name: "Alex"},
		{ID: 2, Name: "Sam"},
	}

	// API

	mux.HandleFunc("/api/services", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		writeJSON(w, services)
	})

	mux.HandleFunc("/api/staff", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		writeJSON(w, staff)
	})

	// CORS
	handler := withCORS(mux)

	log.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.SetIndent("", " ")
	_ = enc.Encode(v)
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PATCH,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

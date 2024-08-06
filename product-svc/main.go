package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ProductCategory struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func main() {
	r := chi.NewRouter()
	r.Get("/api/product-category", func(w http.ResponseWriter, r *http.Request) {
		categories := []ProductCategory{
			{
				ID:   1,
				Name: "Alat Tulis",
			},
			{
				ID:   2,
				Name: "Perlengkapan Rumah Tangga",
			},
			{
				ID:   3,
				Name: "Fashion",
			},
		}

		result, err := json.Marshal(&categories)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Terjadi kesalahan pada server"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(result)
	})

	fmt.Println("your server is running on port 8080")
	http.ListenAndServe(":8080", r)
}

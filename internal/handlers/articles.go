package handlers

import (
    "database/sql"
    "encoding/json"
    "net/http"
)

func ArticlesHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            // Logic for retrieving articles
            articles, err := GetAllArticles(db)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            json.NewEncoder(w).Encode(articles)
        }
        // Implement other methods: POST, PUT, DELETE
    }
}
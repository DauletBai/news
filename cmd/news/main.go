package main

import (
    "log"
    "net/http"
    "https://github.com/DauletBai/news/internal/config"
    "https://github.com/DauletBai/news/internal/database"
    "https://github.com/DauletBai/news/internal/handlers"
    "path/filepath"
    "text/template"
)

func main() {
    cfg, err := config.LoadConfig(".")
    if err != nil {
        log.Fatalf("cannot load config: %v", err)
    }

    conn := database.Connect(cfg.DB)
    defer conn.Close()

    // Setting up routes for static files
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/assets"))))

    // Home page with template rendering
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        templates := []string{
            filepath.Join("static", "tmpl", "home.html"),
            filepath.Join("static", "tmpl", "header.html"),
            filepath.Join("static", "tmpl", "footer.html"),
        }
        tmpl, err := template.ParseFiles(templates...)
        if err != nil {
            http.Error(w, "Unable to load templates", http.StatusInternalServerError)
            log.Println("Error parsing templates:", err)
            return
        }

        // Home page render
        err = tmpl.ExecuteTemplate(w, "home.html", nil)
        if err != nil {
            http.Error(w, "Unable to render page", http.StatusInternalServerError)
            log.Println("Error executing template:", err)
        }
    })

    // Adding handlers CRUD articles, categories and users
    http.HandleFunc("/articles", handlers.ArticlesHandler(conn))
    http.HandleFunc("/categories", handlers.CategoriesHandler(conn))
    http.HandleFunc("/users", handlers.UsersHandler(conn))

    log.Println("Server running on port", cfg.ServerPort)
    log.Fatal(http.ListenAndServe(cfg.ServerPort, nil))
}
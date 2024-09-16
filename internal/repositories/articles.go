package repositories

import (
    "database/sql"
    "news/internal/models"
)

func GetAllArticles(db *sql.DB) ([]models.Article, error) {
    rows, err := db.Query("SELECT id, title, content FROM articles")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var articles []models.Article
    for rows.Next() {
        var article models.Article
        err := rows.Scan(&article.ID, &article.Title, &article.Content)
        if err != nil {
            return nil, err
        }
        articles = append(articles, article)
    }
    return articles, nil
}
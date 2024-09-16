package repositories

import (
    "database/sql"
    "https://github.com/DauletBai/news/internal/models"
)

type ArticleRepository struct {
    db *sql.DB
}

func NewArticleRepository(db *sql.DB) *ArticleRepository {
    return &ArticleRepository{db: db}
}

func (r *ArticleRepository) GetAllArticles() ([]models.Article, error) {
    rows, err := r.db.Query("SELECT id, title, content, author_id, category_id FROM articles")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var articles []models.Article
    for rows.Next() {
        var article models.Article
        if err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.AuthorID, &article.CategoryID); err != nil {
            return nil, err
        }
        articles = append(articles, article)
    }
    return articles, nil
}

func (r *ArticleRepository) GetArticleByID(id int) (*models.Article, error) {
    var article models.Article
    err := r.db.QueryRow("SELECT id, title, content, author_id, category_id FROM articles WHERE id = $1", id).Scan(
        &article.ID, &article.Title, &article.Content, &article.AuthorID, &article.CategoryID)
    if err != nil {
        return nil, err
    }
    return &article, nil
}

func (r *ArticleRepository) CreateArticle(article models.Article) error {
    _, err := r.db.Exec("INSERT INTO articles (title, content, author_id, category_id) VALUES ($1, $2, $3, $4)",
        article.Title, article.Content, article.AuthorID, article.CategoryID)
    return err
}

func (r *ArticleRepository) UpdateArticle(article models.Article) error {
    _, err := r.db.Exec("UPDATE articles SET title=$1, content=$2, author_id=$3, category_id=$4 WHERE id=$5",
        article.Title, article.Content, article.AuthorID, article.CategoryID, article.ID)
    return err
}

func (r *ArticleRepository) DeleteArticle(id int) error {
    _, err := r.db.Exec("DELETE FROM articles WHERE id=$1", id)
    return err
}

package services

import (
    "database/sql"
    "https://github.com/DauletBai/news/internal/models"
    "https://github.com/DauletBai/news/internal/repositories"
)

type ArticleService struct {
    repo *repositories.ArticleRepository
}

func NewArticleService(db *sql.DB) *ArticleService {
    return &ArticleService{
        repo: repositories.NewArticleRepository(db),
    }
}

func (s *ArticleService) GetAllArticles() ([]models.Article, error) {
    return s.repo.GetAllArticles()
}

func (s *ArticleService) GetArticleByID(id int) (*models.Article, error) {
    return s.repo.GetArticleByID(id)
}

func (s *ArticleService) CreateArticle(article models.Article) error {
    return s.repo.CreateArticle(article)
}

func (s *ArticleService) UpdateArticle(article models.Article) error {
    return s.repo.UpdateArticle(article)
}

func (s *ArticleService) DeleteArticle(id int) error {
    return s.repo.DeleteArticle(id)
}
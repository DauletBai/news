package services

import (
    "database/sql"
    "https://github.com/DauletBai/news/internal/models"
    "https://github.com/DauletBai/news/internal/repositories"
)

type UserService struct {
    repo *repositories.UserRepository
}

func NewUserService(db *sql.DB) *UserService {
    return &UserService{
        repo: repositories.NewUserRepository(db),
    }
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
    return s.repo.GetAllUsers()
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
    return s.repo.GetUserByID(id)
}

func (s *UserService) CreateUser(user models.User) error {
    return s.repo.CreateUser(user)
}

func (s *UserService) UpdateUser(user models.User) error {
    return s.repo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id int) error {
    return s.repo.DeleteUser(id)
}

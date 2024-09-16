package repositories

import (
    "database/sql"
    "https://github.com/DauletBai/news/internal/models"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
    rows, err := r.db.Query("SELECT id, name, role, email FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var user models.User
        if err := rows.Scan(&user.ID, &user.Name, &user.Role, &user.Email); err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    return users, nil
}

func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
    var user models.User
    err := r.db.QueryRow("SELECT id, name, role, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Role, &user.Email)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepository) CreateUser(user models.User) error {
    _, err := r.db.Exec("INSERT INTO users (name, role, email, hash_password) VALUES ($1, $2, $3, $4)",
        user.Name, user.Role, user.Email, user.HashPassword)
    return err
}

func (r *UserRepository) UpdateUser(user models.User) error {
    _, err := r.db.Exec("UPDATE users SET name=$1, role=$2, email=$3 WHERE id=$4",
        user.Name, user.Role, user.Email, user.ID)
    return err
}

func (r *UserRepository) DeleteUser(id int) error {
    _, err := r.db.Exec("DELETE FROM users WHERE id=$1", id)
    return err
}
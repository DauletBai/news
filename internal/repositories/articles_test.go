package repositories

import (
    "database/sql"
    "testing"

    "github.com/DATA-DOG/go-sqlmock"
    "https://github.com/DauletBai/news/internal/models"
)

func TestGetAllArticles(t *testing.T) {
    // Создаем mock базы данных
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("error creating mock DB: %v", err)
    }
    defer db.Close()

    // Определяем ожидаемый результат запроса
    rows := sqlmock.NewRows([]string{"id", "title", "content"}).
        AddRow(1, "Test Article", "This is a test article.")

    mock.ExpectQuery("SELECT id, title, content FROM articles").
        WillReturnRows(rows)

    // Вызываем тестируемую функцию
    articles, err := GetAllArticles(db)
    if err != nil {
        t.Errorf("error calling GetAllArticles: %v", err)
    }

    // Проверка результата
    if len(articles) != 1 {
        t.Errorf("expected 1 article, got %v", len(articles))
    }

    expected := models.Article{ID: 1, Title: "Test Article", Content: "This is a test article."}
    if articles[0] != expected {
        t.Errorf("expected %v, got %v", expected, articles[0])
    }

    // Проверяем, что все ожидания SQL-запросов были выполнены
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expectations: %v", err)
    }
}
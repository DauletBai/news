package handlers

import (
    "database/sql"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestArticlesHandler(t *testing.T) {
    // Подготовка тестовой базы данных
    db, _, err := sqlmock.New()
    if err != nil {
        t.Fatalf("failed to open mock sql db, got error: %v", err)
    }
    defer db.Close()

    req, err := http.NewRequest("GET", "/articles", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Записываем ответ в тестовый ResponseRecorder
    rr := httptest.NewRecorder()
    handler := ArticlesHandler(db)

    // Вызов обработчика
    handler.ServeHTTP(rr, req)

    // Проверяем код статуса
    if rr.Code != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
    }

    // Проверка на ожидаемое содержание ответа
    expected := `[{"id":1,"title":"Test Article","content":"This is a test article."}]`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
}
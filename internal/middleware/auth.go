package middleware

import (
    "net/http"
    "strings"
)

func Auth(next http.HandlerFunc, roles ...string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // Имитация проверки токена и получения роли пользователя
        userRole := parseToken(token)
        if userRole == "" {
            http.Error(w, "Forbidden", http.StatusForbidden)
            return
        }

        // Проверка, имеет ли пользователь доступ к роли
        for _, role := range roles {
            if strings.EqualFold(userRole, role) {
                next.ServeHTTP(w, r)
                return
            }
        }

        http.Error(w, "Forbidden", http.StatusForbidden)
    }
}

func parseToken(token string) string {
    // Простая имитация проверки токена
    if token == "admin-token" {
        return "admin"
    } else if token == "editor-token" {
        return "editor"
    } else if token == "user-token" {
        return "user"
    }
    return ""
}
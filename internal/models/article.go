package models

type Article struct {
    ID      int    `json:"id"`
    Title   string `json:"title"`
    Content string `json:"content"`
    AuthorID int   `json:"author_id"`
    CategoryID int `json:"category_id"`
}
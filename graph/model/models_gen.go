// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Article struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  bool   `json:"status"`
	User    *User  `json:"user"`
}

type NewArticle struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  string `json:"userId"`
}

type NewUser struct {
	Name string `json:"name"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

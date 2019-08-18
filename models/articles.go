package models

type Article struct {
	Id      uint64	`json:"id"`
	Title   string	`valid:"required" json"title"`
	Content string	`valid:"required" json"content"`
	Author  string	`valid:"required" json:"author"`
}
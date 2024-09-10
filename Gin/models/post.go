package models

import "github.com/uptrace/bun"

type Post struct {
	bun.BaseModel `bun:",alias:posts"`
	ID            uint64 `bun:"id,pk,autoincrement,notnull" json:"id"`
	Title         string `bun:"title,notnull" json:"title"`
	Content       string `bun:"content,notnull" json:"content"`
	Image         string `bun:"image,null" json:"image"`
	UserID        uint64 `bun:"user_id,notnull" json:"user_id"`
	User          *User  `bun:"rel:belongs-to,join:user_id=id"`
}

// User:
// 	has_many :posts

// Post:
// 	has_many :comments
// 	belongs_to :user

// Comment:
// 	belongs_to :post
// 	belongs_to :user

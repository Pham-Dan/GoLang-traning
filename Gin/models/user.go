package models

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:",alias:users"`

	ID       uint64  `bun:"id,pk,autoincrement" json:"id"`
	Name     string  `bun:"name,notnull" json:"name"`
	Email    string  `bun:"email,notnull,unique" json:"email"`
	Password string  `bun:"password,notnull" json:"-"`
	Post     []*Post `bun:"rel:has-many,join:id=user_id"`
}

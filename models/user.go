package models

import "time"

type User struct {
	ID		  string	`json:"id"`
	Username  string    `json:"username" validate:"required,min=5,max=15"`
	Name      string    `json:"name" validate:"required,min=5,max=50"`
	Password  string    `json:"password" validate:"required,min=5,max=15"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"Updated_at"`
}

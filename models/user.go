package models

import "time"

type User struct {
<<<<<<< HEAD
	ID        uint      `json:"id" validate:"required"`
=======
	ID		  string	`json:"id"`
>>>>>>> 6d8726bd4d290768c83b7527e52e7c883f358143
	Username  string    `json:"username" validate:"required,min=5,max=15"`
	Name      string    `json:"name" validate:"required,min=5,max=50"`
	Password  string    `json:"password" validate:"required,min=5,max=15"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"Updated_at"`
}

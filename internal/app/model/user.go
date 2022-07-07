package model

import "time"

type (
	User struct {
		ID        string
		Name      string
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt *time.Time
	}

	UserUsecase interface {
	}

	UserRepository interface {
	}
)

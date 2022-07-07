package usecase

import "serenade/internal/app/model"

type implUser struct {
}

func NewUserUsecase() model.UserUsecase {
	return &implUser{}
}

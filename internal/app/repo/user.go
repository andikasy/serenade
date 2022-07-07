package repo

import "serenade/internal/app/model"

type implUser struct {
}

func NewUserRepository() model.UserRepository {
	return &implUser{}
}

func (i *implUser) Insert() error {
	return nil
}

func (i *implUser) SelectOne() (*model.User, error) {
	return nil, nil
}

func (i *implUser) SelectAll() ([]*model.User, error) {
	return nil, nil
}

func (i *implUser) Update() error {
	return nil
}

func (i *implUser) Delete() error {
	return nil
}

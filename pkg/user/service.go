package user

import "ocra/pkg/entities"

type Service interface {
	Login(email, password string) (*entities.User, error)
}

type service struct {
	repository Repository
}

func (s *service) Login(email string, password string) (*entities.User, error) {
	return s.repository.Login(email, password)
}

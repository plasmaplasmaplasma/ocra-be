package user

import "ocra/pkg/entities"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Service interface {
	Login(email, password string) (*entities.User, error)
}

type service struct {
	repository Repository
}

func (s *service) Login(email string, password string) (*entities.User, error) {
	return s.repository.Login(email, password)
}

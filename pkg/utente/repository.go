package user

import (
	"errors"
	"ocra/pkg/entities"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Repository interface {
	Login(email string, password string) (*entities.User, error)
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		DB: db,
	}
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (r *repository) Login(email, password string) (*entities.User, error) {
	var user entities.User

	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	return &entities.User{
		BaseEntity: entities.BaseEntity{ID: user.ID},
		Email:      user.Email,
		Username:   user.Username,
	}, nil
}

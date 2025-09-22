package user

import "my_server/domain"

type UserRepo interface {
	Create(user domain.User) (*domain.User, error)
	Find(email, password string) (*domain.User, error)
	Get(userId int) (*domain.User, error)
	List() ([]*domain.User, error)
	Delete(userId int) error
	Update(user domain.User) (*domain.User, error)
}

package store

import (
	"github.com/PiterPentester/goRESTLearn/internal/app/model"
)

type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}

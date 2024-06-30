package processor

import (
	"errors"
	"some-httpserver/internal/app/db"
	"some-httpserver/internal/app/models"
)

type UsersProcessor struct {
	storage *db.UserStorage
}

func NewUsersProcessor(storage *db.UserStorage) *UsersProcessor {
	processor := &UsersProcessor{
		storage: storage,
	}
	return processor
}

func (processor *UsersProcessor) CreateUser(user models.User) error {
	if user.Name == "" {
		return errors.New("name should not be empty")
	}
	return processor.storage.CreateUser(user)
}

func (processor *UsersProcessor) FindUser(id uint64) (models.User, error) {
	user := processor.storage.GetUserById(id)

	if user.Id != id {
		return user, errors.New("user not found")
	}
	return user, nil
}

func (processor *UsersProcessor) ListUsers(nameFilter string) ([]models.User, error) {
	return processor.storage.GetUsersList(nameFilter), nil
}

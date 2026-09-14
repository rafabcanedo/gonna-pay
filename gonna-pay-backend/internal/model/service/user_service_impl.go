package service

import (
	"errors"

	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/configuration/logger"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/repository"
)

type UserService interface {
	Create(user *domains.User) (*domains.User, error)
	FindAll(page, limit int) ([]*domains.User, int64, error)
	FindByID(id string) (*domains.User, error)
	FindByEmail(email string) (*domains.User, error)
	Update(user *domains.User) (*domains.User, error)
	Delete(id string) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Create(user *domains.User) (*domains.User, error) {
	existing, err := s.repo.FindByEmail(user.Email)
	if err != nil && !errors.Is(err, domains.ErrNotFound) {
		logger.Error("error checking email uniqueness", err)
		return nil, err
	}
	if existing != nil {
		return nil, domains.NewConflictError("email already in use")
	}

	if err := user.EncryptPassword(); err != nil {
		logger.Error("error encrypting password", err)
		return nil, err
	}

	created, err := s.repo.Create(user)
	if err != nil {
		logger.Error("error creating user", err)
		return nil, err
	}

	return created, nil
}

func (s *userService) FindAll(page, limit int) ([]*domains.User, int64, error) {
	offset := (page - 1) * limit

	users, total, err := s.repo.FindAll(limit, offset)
	if err != nil {
		logger.Error("error finding all users", err)
		return nil, 0, err
	}

	return users, total, nil
}

func (s *userService) FindByID(id string) (*domains.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		if !errors.Is(err, domains.ErrNotFound) {
			logger.Error("error finding user by id", err)
		}
		return nil, err
	}

	return user, nil
}

func (s *userService) FindByEmail(email string) (*domains.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		if !errors.Is(err, domains.ErrNotFound) {
			logger.Error("error finding user by email", err)
		}
		return nil, err
	}

	return user, nil
}

func (s *userService) Update(user *domains.User) (*domains.User, error) {
	if _, err := s.repo.FindByID(user.ID); err != nil {
		if !errors.Is(err, domains.ErrNotFound) {
			logger.Error("error finding user on update", err)
		}
		return nil, err
	}

	if user.Email != "" {
		existing, err := s.repo.FindByEmail(user.Email)
		if err != nil && !errors.Is(err, domains.ErrNotFound) {
			logger.Error("error checking email uniqueness on update", err)
			return nil, err
		}
		if existing != nil && existing.ID != user.ID {
			return nil, domains.NewConflictError("email already in use")
		}
	}

	if user.Password != "" {
		if err := user.EncryptPassword(); err != nil {
			logger.Error("error encrypting password on update", err)
			return nil, err
		}
	}

	updated, err := s.repo.Update(user)
	if err != nil {
		logger.Error("error updating user", err)
		return nil, err
	}

	return updated, nil
}

func (s *userService) Delete(id string) error {
	if _, err := s.repo.FindByID(id); err != nil {
		if !errors.Is(err, domains.ErrNotFound) {
			logger.Error("error finding user on delete", err)
		}
		return err
	}

	if err := s.repo.Delete(id); err != nil {
		logger.Error("error deleting user", err)
		return err
	}

	return nil
}

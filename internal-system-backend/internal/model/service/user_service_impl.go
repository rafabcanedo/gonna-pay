package service

import (
	"errors"

	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/configuration/logger"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/domains"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/repository"
)

type UserService interface {
	Create(domain domains.UserDomainInterface) (domains.UserDomainInterface, error)
	FindAll() ([]domains.UserDomainInterface, error)
	FindByID(id string) (domains.UserDomainInterface, error)
	FindByEmail(email string) (domains.UserDomainInterface, error)
	Update(domain domains.UserDomainInterface) (domains.UserDomainInterface, error)
	Delete(id string) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Create(domain domains.UserDomainInterface) (domains.UserDomainInterface, error) {
	existing, err := s.repo.FindByEmail(domain.GetEmail())
	if err != nil && !errors.Is(err, domains.ErrNotFound) {
		logger.Error("error checking email uniqueness", err)
		return nil, err
	}
	if existing != nil {
		return nil, domains.NewConflictError("email already in use")
	}

	if err := domain.EncryptPassword(); err != nil {
		logger.Error("error encrypting password", err)
		return nil, err
	}

	created, err := s.repo.Create(domain)
	if err != nil {
		logger.Error("error creating user", err)
		return nil, err
	}

	return created, nil
}

func (s *userService) FindAll() ([]domains.UserDomainInterface, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		logger.Error("error finding all users", err)
		return nil, err
	}

	return users, nil
}

func (s *userService) FindByID(id string) (domains.UserDomainInterface, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		if !errors.Is(err, domains.ErrNotFound) {
			logger.Error("error finding user by id", err)
		}
		return nil, err
	}

	return user, nil
}

func (s *userService) FindByEmail(email string) (domains.UserDomainInterface, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		if !errors.Is(err, domains.ErrNotFound) {
			logger.Error("error finding user by email", err)
		}
		return nil, err
	}

	return user, nil
}

func (s *userService) Update(domain domains.UserDomainInterface) (domains.UserDomainInterface, error) {
	if domain.GetPassword() != "" {
		if err := domain.EncryptPassword(); err != nil {
			logger.Error("error encrypting password on update", err)
			return nil, err
		}
	}

	updated, err := s.repo.Update(domain)
	if err != nil {
		logger.Error("error updating user", err)
		return nil, err
	}

	return updated, nil
}

func (s *userService) Delete(id string) error {
	if err := s.repo.Delete(id); err != nil {
		logger.Error("error deleting user", err)
		return err
	}

	return nil
}

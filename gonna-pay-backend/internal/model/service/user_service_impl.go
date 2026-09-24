package service

import (
	"context"
	"errors"
	"time"

	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/auth"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/configuration/logger"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/repository"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/repository/entity/enums"
)

type UserService interface {
	Create(ctx context.Context, user *domains.User) (*domains.User, error)
	FindAll(ctx context.Context, page, limit int) ([]*domains.User, int64, error)
	FindByID(ctx context.Context, id string) (*domains.User, error)
	FindByEmail(ctx context.Context, email string) (*domains.User, error)
	Update(ctx context.Context, user *domains.User) (*domains.User, error)
	Delete(ctx context.Context, id string) error
}

type userService struct {
	repo repository.UserRepository
	emailTokenRepo repository.EmailTokenRepository
	emailSvc EmailService
}

func NewUserService(repo repository.UserRepository, emailTokenRepo repository.EmailTokenRepository, emailSvc EmailService) UserService {
	return &userService{repo: repo, emailTokenRepo: emailTokenRepo, emailSvc: emailSvc}
}

func (s *userService) Create(ctx context.Context, user *domains.User) (*domains.User, error) {
	existing, err := s.repo.FindByEmail(ctx, user.Email)
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

	created, err := s.repo.Create(ctx, user)
	if err != nil {
		logger.Error("error creating user", err)
		return nil, err
	}

	token, err := auth.GenerateRefreshToken()
	if err != nil {
		logger.Error("error generating verification token", err)
		return created, nil
	}

	tokenHash := auth.HashToken(token)
	expiresAt := time.Now().Add(24 * time.Hour)

	if err := s.emailTokenRepo.Save(ctx, created.ID, tokenHash, enums.EmailTokenTypeVerification, expiresAt); err != nil {
		logger.Error("error saving verification token", err)
		return created, nil
	}

	if err := s.emailSvc.SendVerificationEmail(ctx, created.Email, created.Name, token); err != nil {
		logger.Error("error sending verification email", err)
	}

	return created, nil
}

func (s *userService) FindAll(ctx context.Context, page, limit int) ([]*domains.User, int64, error) {
	offset := (page - 1) * limit

	users, total, err := s.repo.FindAll(ctx, limit, offset)
	if err != nil {
		logger.Error("error finding all users", err)
		return nil, 0, err
	}

	return users, total, nil
}

func (s *userService) FindByID(ctx context.Context, id string) (*domains.User, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if !errors.Is(err, domains.ErrNotFound) {
			logger.Error("error finding user by id", err)
		}
		return nil, err
	}

	return user, nil
}

func (s *userService) FindByEmail(ctx context.Context, email string) (*domains.User, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		if !errors.Is(err, domains.ErrNotFound) {
			logger.Error("error finding user by email", err)
		}
		return nil, err
	}

	return user, nil
}

func (s *userService) Update(ctx context.Context, user *domains.User) (*domains.User, error) {
	if _, err := s.repo.FindByID(ctx, user.ID); err != nil {
		if !errors.Is(err, domains.ErrNotFound) {
			logger.Error("error finding user on update", err)
		}
		return nil, err
	}

	if user.Email != "" {
		existing, err := s.repo.FindByEmail(ctx, user.Email)
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

	updated, err := s.repo.Update(ctx, user)
	if err != nil {
		logger.Error("error updating user", err)
		return nil, err
	}

	return updated, nil
}

func (s *userService) Delete(ctx context.Context, id string) error {
	if _, err := s.repo.FindByID(ctx, id); err != nil {
		if !errors.Is(err, domains.ErrNotFound) {
			logger.Error("error finding user on delete", err)
		}
		return err
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		logger.Error("error deleting user", err)
		return err
	}

	return nil
}

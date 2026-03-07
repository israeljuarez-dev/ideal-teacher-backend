package service

import (
	"context"

	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/adapters/dto/user"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/adapters/mapper"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/core/domain"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/core/ports"

)

type userService struct {
	repository ports.UserRepository
}

func NewUserService(repo ports.UserRepository) ports.UserService {
	return &userService{
		repository: repo,
	}
}

func (s *userService) GetByID(ctx context.Context, ID int32) (*user.UserResponse, error) {
	u, err := s.repository.GetByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	userResponse := mapper.UserToResponse(u)

	return userResponse, nil
}

func (s *userService) GetByEmail(ctx context.Context, email string) (*user.UserResponse, error) {
	u, err := s.repository.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	userResponse := mapper.UserToResponse(u)

	return userResponse, nil
}

func (s *userService) GetAll(ctx context.Context, limit, offset int32) (*user.UsersListResponse, error) {
	users, err := s.repository.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	usersResponse := mapper.UsersToListResponse(users)

	return usersResponse, nil
}

func (s *userService) Create(ctx context.Context, ur *user.CreateUserRequest) (*user.UserResponse, error) {
	hashedPassword, err := HashPassword(ur.Password)
    if err != nil {
        return nil, err 
    }

	u := &domain.User{
		Email:    ur.Email,
		Password: hashedPassword, 
		FullName: ur.FullName,
		Role:     domain.RoleStudent,
	}

	savedUser, err := s.repository.Create(ctx, u)
	if err != nil {
		return nil, err
	}

	userResponse := mapper.UserToResponse(savedUser)

	return userResponse, nil
}

func (s *userService) Update(ctx context.Context, ID int32, ur *user.UpdateUserRequest) (*user.UserResponse, error) {
    existingUser, err := s.repository.GetByID(ctx, ID)
    if err != nil {
        return nil, err
    }

    existingUser.FullName = *ur.FullName
    existingUser.Email = *ur.Email

    updatedUser, err := s.repository.Update(ctx, existingUser)
    if err != nil {
        return nil, err
    }

	userResponse := mapper.UserToResponse(updatedUser)

    return  userResponse, nil
}

func (s *userService) Delete(ctx context.Context, ID int32) error {
	if err := s.repository.Delete(ctx, ID); err != nil {
		return err
	}

	return nil
}
package ports

import (
	"context"
	"net/http"

	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/adapters/dto/user"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/core/domain"

)

type UserRepository interface {
	GetByID(ctx context.Context, ID int32) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	GetAll(ctx context.Context, limit, offset int32) (domain.Users, error)
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) (*domain.User, error)
	Delete(ctx context.Context, ID int32) error
}

type UserService interface {
	GetByID(ctx context.Context, ID int32) (*user.UserResponse, error)
	GetByEmail(ctx context.Context, email string) (*user.UserResponse, error)
	GetAll(ctx context.Context, limit, offset int32) (*user.UsersListResponse, error)
	Create(ctx context.Context, user *user.CreateUserRequest) (*user.UserResponse, error)
	Update(ctx context.Context, user *user.UpdateUserRequest) (*user.UserResponse, error)
	Delete(ctx context.Context, ID int32) error
}

type UserHandler interface {
	GetByID(w http.ResponseWriter, r *http.Request)
	GetByEmail(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
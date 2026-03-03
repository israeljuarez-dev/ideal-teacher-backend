package repository

import (
	"context"

	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/adapters/storage/postgres"
	pg "github.com/israeljuarez-dev/ideal-teacher-backend/internal/adapters/storage/postgres/db"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/core/domain"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/core/ports"

)

type userRepository struct {
	db *postgres.DB
	query *pg.Queries
} 

func NewUserRepository(db *postgres.DB) ports.UserRepository {
    return &userRepository{
        db: db,
        query:  pg.New(db.Pool),
    }
}

func (r *userRepository) GetByID(ctx context.Context, id int32) (*domain.User, error) {
	u, err := r.query.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:       u.ID,
		Email:    u.Email,
		Password: u.Password,
		FullName: u.FullName,
		Role:     domain.Role(u.Role),
	}, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	u, err := r.query.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:       u.ID,
		Email:    u.Email,
		Password: u.Password,
		FullName: u.FullName,
		Role:     domain.Role(u.Role),
	}, nil
}

func (r *userRepository) GetAll(ctx context.Context, limit, offset int32) (domain.Users, error) {
	rows, err := r.query.GetAll(ctx, pg.GetAllParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	var users domain.Users
	for _, row := range rows {
		users = append(users, domain.User{
			ID:       row.ID,
			Email:    row.Email,
			FullName: row.FullName,
			Role:     domain.Role(row.Role),
		})
	}

	return users, nil
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	u, err := r.query.CreateUser(ctx, pg.CreateUserParams{
		Email:    user.Email,
		Password: user.Password,
		FullName: user.FullName,
	})
	if err != nil {
		return nil, err
	}

	user.ID = u.ID

	return user, nil
}

func (r *userRepository) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
	err := r.query.UpdateUser(ctx, pg.UpdateUserParams{
		ID:       user.ID,
		Email:    user.Email,
		FullName: user.FullName,
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) Delete(ctx context.Context, id int32) error {
	if err := r.query.DeleteUser(ctx, id); err != nil{
		return err
	}

	return nil
}
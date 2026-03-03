package mapper

import (
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/adapters/dto/user"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/core/domain"

)

// CreateRequestToUser convierte CreateUserRequest en domain.User
// Nota: El password debe hashearse en el service.
func CreateRequestToUser(r *user.CreateUserRequest) *domain.User {
	return &domain.User{
		Email:    r.Email,
		Password: r.Password,
		FullName: r.FullName,
	}
}

// UserToResponse convierte domain.User en UserResponse
func UserToResponse(u *domain.User) *user.UserResponse {
	if u == nil {
		return nil
	}

	return &user.UserResponse{
		ID:       int32(u.ID),
		Email:    u.Email,
		FullName: u.FullName,
		Role:     string(u.Role),
		Status:   string(u.Status),
	}
}

// UsersToListResponse convierte un slice de usuarios en UsersListResponse
func UsersToListResponse(users domain.Users, total int64) *user.UsersListResponse {
	data := make([]user.UserResponse, 0, len(users))

	for _, u := range users {
		data = append(data, user.UserResponse{
			ID:       int32(u.ID),
			Email:    u.Email,
			FullName: u.FullName,
			Role:     string(u.Role),
			Status:   string(u.Status),
		})
	}

	return &user.UsersListResponse{
		Data:  data,
		Total: total,
	}
}

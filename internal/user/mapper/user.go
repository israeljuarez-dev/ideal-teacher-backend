package mapper

import (
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/domain"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/dto"
)

// toUserResponse — mapper de domain.User → dto.UserResponse
func ToUserResponse(u *domain.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:       u.ID,
		Email:    u.Email,
		FullName: u.FirstName + " " + u.LastName,
		Status:   string(u.Status),
	}
}

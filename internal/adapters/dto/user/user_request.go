package user

type CreateUserRequest struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=8,max=72"`
	FullName string `json:"full_name" validate:"required,min=3,max=200"`
	Role     string `json:"role" validate:"required"`
}

type UpdateUserRequest struct {
	Email    *string `json:"email,omitempty" validate:"omitempty,email,max=255"`
	FullName *string `json:"full_name,omitempty" validate:"omitempty,min=3,max=200"`
}
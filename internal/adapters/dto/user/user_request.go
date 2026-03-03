package user

type CreateUserRequest struct {
	Email    string `json:"email" validate:"required,email,max=150"`
	Password string `json:"password" validate:"required,min=8,max=72"`
	FullName string `json:"full_name" validate:"required,min=3,max=200"`
}

type UpdateUserRequest struct {
	Email    *string `json:"email,omitempty" validate:"omitempty,email,max=255"`
	FullName *string `json:"full_name,omitempty" validate:"omitempty,min=3,max=200"`
	Status   *string `json:"status,omitempty" validate:"omitempty,oneof=active inactive"`
}
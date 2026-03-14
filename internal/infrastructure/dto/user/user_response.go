package user

type UserResponse struct {
	ID        int32     `json:"id"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	Role      string    `json:"role"`
	Status    string    `json:"status"`
}

type UsersListResponse struct {
	Data  []UserResponse `json:"data"`
	Total int          `json:"total"`
}
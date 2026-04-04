package pipes

type (
	CreateUserIn struct {
		Email     string `json:"email" validate:"required,email,max=255"`
		Password  string `json:"password" validate:"required,min=8,max=72"`
		FirstName string `json:"first_name" validate:"required,min=3,max=200"`
		LastName  string `json:"last_name" validate:"required,min=3,max=200"`
	}

	UpdateUserIn struct {
		Email     *string `json:"email,omitempty" validate:"omitempty,email,max=255"`
		FirstName *string `json:"first_name,omitempty" validate:"omitempty,min=3,max=200"`
		LastName  *string `json:"last_name,omitempty" validate:"omitempty,min=3,max=200"`
	}
)

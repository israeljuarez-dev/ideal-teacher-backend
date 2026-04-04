package service

type (
	LoginIn struct {
		Email    string
		Password string
	}

	LoginOut struct {
		Token     string
		ExpiresIn int
	}
)

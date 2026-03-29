package models

import "github.com/jackc/pgx/v5/pgtype"

type (
	InsertUserParams struct {
		ID        pgtype.UUID
		Email     string
		Password  string
		FirstName string
		LastName  string
	}
)

package service

import (
	"context"
	"fmt"
)

func (s *Service) GetByEmail(ctx context.Context, email string) (*GetByEmailUserOut, error) {
	udb, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	u := &GetByEmailUserOut{
		ID:        udb.ID,
		Email:     udb.Email,
		FirstName: udb.FirstName,
		LastName:  udb.LastName,
		Status:    udb.Status,
		CreatedAt: udb.CreatedAt,
		RoleName:  udb.RoleName,
	}

	return u, nil
}

package service

import (
	"context"
	"fmt"

	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/auth/jwt"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/auth/myerrors"
)

func (s *service) Login(ctx context.Context, in *LoginIn) (*LoginOut, error) {
	u, err := s.repo.GetByEmail(ctx, in.Email)
	if err != nil {
		s.log.Error("service.Login: user not found", "email", in.Email)
		return nil, &myerrors.AuthError{
			Msg: "invalid user or password",
			Err: myerrors.InvalidEmailOrPassword,
		}
	}

	if err := comparePassword(in.Password, u.Password); err != nil {
		s.log.Error("service.Login: wrong password", "email", in.Email)
		return nil, &myerrors.AuthError{
			Msg: "invalid user or password",
			Err: myerrors.InvalidEmailOrPassword, 
		}
	}

	accessToken, err := jwt.GenerateToken(u, s.cfg)
	if err != nil {
		 s.log.Error("service.Login: token generation failed", "error", err)
        return nil, fmt.Errorf("error generating token: %w", err)
	}

	response := &LoginOut{
		Token:     accessToken,
		ExpiresIn: int(s.cfg.ExpirationTime),
	}

	return response, nil
}

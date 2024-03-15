package service

import (
	"auth-service/pkg/config"
	"auth-service/pkg/jwt"
	"auth-service/pkg/models"
	"auth-service/pkg/pb"
	"context"
	"database/sql"
)

type UserServices interface {
	UserService
}

type userServices struct {
	UserService
}

type UserService interface {
	GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error)
	GetCurrentUser(ctx context.Context, req *pb.GetCurrentUserRequest) (*pb.GetUserResponse, error)
	DeleteCurrentUser(ctx context.Context, req *pb.DeleteCurrentUserRequest) (*pb.EmptyResponse, error)
}

type userService struct {
	db  *sql.DB
	cfg *config.Config
}

func NewUserService(db *sql.DB, cfg *config.Config) UserService {
	return &userServices{
		UserService: &userService{
			db:  db,
			cfg: cfg,
		},
	}
}

func (u *userService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var user models.User
	err := u.db.QueryRowContext(ctx, "SELECT * FROM users WHERE id = $1", req.GetId()).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Contacts, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		Id:       user.ID.String(),
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (u *userService) GetCurrentUser(ctx context.Context, req *pb.GetCurrentUserRequest) (*pb.GetUserResponse, error) {
	wrapper := jwt.JwtWrapper{
		SecretKey:       u.cfg.JWT_SECRET_KEY,
		ExpirationHours: u.cfg.JWT_EXPIRATION_HOURS,
	}
	userId, err := wrapper.ValidateToken(req.GetJwt())

	if err != nil {
		return nil, err
	}

	var user models.User
	err = u.db.QueryRowContext(ctx, "SELECT * FROM users WHERE id = $1", userId).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Contacts, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		Id:       user.ID.String(),
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (u *userService) DeleteCurrentUser(ctx context.Context, req *pb.DeleteCurrentUserRequest) (*pb.EmptyResponse, error) {
	wrapper := jwt.JwtWrapper{
		SecretKey:       u.cfg.JWT_SECRET_KEY,
		ExpirationHours: u.cfg.JWT_EXPIRATION_HOURS,
	}
	userId, err := wrapper.ValidateToken(req.GetJwt())

	if err != nil {
		return nil, err
	}

	_, err = u.db.ExecContext(ctx, "DELETE FROM users WHERE id = $1", userId)
	if err != nil {
		return nil, err
	}

	return &pb.EmptyResponse{}, nil
}

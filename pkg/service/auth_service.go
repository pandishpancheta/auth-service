package service

import (
	"auth-service/pkg/config"
	"auth-service/pkg/jwt"
	"auth-service/pkg/models"
	"auth-service/pkg/pb"
	"context"
	"database/sql"
	"errors"
	"github.com/gofrs/uuid/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthServices interface {
	AuthService
}

type authServices struct {
	AuthService
}

type AuthService interface {
	Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error)
	Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error)
	ValidateToken(context.Context, *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error)
}

type authService struct {
	db  *sql.DB
	cfg *config.Config
}

func NewAuthService(db *sql.DB, cfg *config.Config) AuthService {
	return &authServices{
		AuthService: &authService{
			db:  db,
			cfg: cfg,
		},
	}
}

func (a *authService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	var user models.User
	err := a.db.QueryRowContext(ctx, "SELECT * FROM users WHERE email = $1", req.GetEmail()).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.GetPassword()))

	if err != nil {
		return nil, err
	}

	wrapper := jwt.JwtWrapper{
		a.cfg.JWT_SECRET_KEY,
		a.cfg.JWT_EXPIRATION_HOURS,
	}
	token, err := wrapper.GenerateToken(user.ID.String(), user.Email)

	if err != nil {
		return nil, err
	}

	return &pb.AuthResponse{
		Jwt: token,
	}, nil
}

func (a *authService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	var user models.User
	var err error
	user.ID, err = uuid.NewV4()
	if err != nil {
		return nil, err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user = models.User{
		ID:        user.ID,
		Username:  req.GetUsername(),
		Email:     req.GetEmail(),
		Password:  string(hashedPassword),
		Contacts:  models.Contacts{},
		CreatedAt: time.Now(),
	}

	_, err = a.db.ExecContext(ctx, "INSERT INTO users (id, username, email, password, created_at) VALUES ($1, $2, $3, $4, $5)", user.ID, user.Username, user.Email, user.Password, user.CreatedAt)
	if err != nil {
		return nil, err
	}

	wrapper := jwt.JwtWrapper{
		a.cfg.JWT_SECRET_KEY,
		a.cfg.JWT_EXPIRATION_HOURS,
	}
	token, err := wrapper.GenerateToken(user.ID.String(), user.Email)

	if err != nil {
		return nil, err
	}

	return &pb.AuthResponse{
		Jwt: token,
	}, nil
}

func (a *authService) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	wrapper := jwt.JwtWrapper{
		SecretKey:       a.cfg.JWT_SECRET_KEY,
		ExpirationHours: a.cfg.JWT_EXPIRATION_HOURS,
	}

	uuid, err := wrapper.ValidateToken(req.GetJwt())
	if err != nil {
		return nil, err
	}

	if !UserExists(a.db, *uuid) {
		return nil, errors.New("user does not exist")
	}

	return &pb.ValidateTokenResponse{
		UserId: uuid.String(),
	}, nil
}

func UserExists(db *sql.DB, uuid uuid.UUID) bool {
	var id string
	err := db.QueryRowContext(context.Background(), "SELECT id FROM users WHERE id = $1", uuid.String()).Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		println(err.Error())
		return false
	} else if err != nil {
		println(err.Error())
		return false
	}
	return true
}

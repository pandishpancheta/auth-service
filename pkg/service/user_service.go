package service

import (
	"auth-service/pkg/config"
	"auth-service/pkg/jwt"
	"auth-service/pkg/models"
	"auth-service/pkg/pb"
	"context"
	"database/sql"
	"github.com/gofrs/uuid/v5"
	"log"
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
	CreateContact(ctx context.Context, req *pb.CreateContactRequest) (*pb.Contact, error)
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
	query := `
        SELECT u.id, u.username, u.email, u.password, u.created_at, c.id, c.email, c.phone, c.instagram, c.other
        FROM users u
        LEFT JOIN contacts c ON u.id = c.user_id
        WHERE u.id = $1
    `

	var user models.User
	log.Println(req.GetId())
	err := u.db.QueryRowContext(ctx, query, req.GetId()).Scan(
		&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt,
		&user.Contacts.ID, &user.Contacts.Email, &user.Contacts.Phone, &user.Contacts.Instagram, &user.Contacts.Other,
	)
	if err != nil {
		return nil, err
	}
	log.Println("Fetched from db")
	return &pb.GetUserResponse{
		Id:       user.ID.String(),
		Username: user.Username,
		Email:    user.Email,
		Contact: &pb.Contact{
			Id:        user.Contacts.ID.String(),
			Email:     user.Contacts.Email.String,
			Phone:     user.Contacts.Phone.String,
			Instagram: user.Contacts.Instagram.String,
			Other:     user.Contacts.Other.String,
		},
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

	query := `
        SELECT u.id, u.username, u.email, u.password, u.created_at, c.id, c.email, c.phone, c.instagram, c.other
        FROM users u
        LEFT JOIN contacts c ON u.id = c.user_id
        WHERE u.id = $1
    `

	var user models.User
	err = u.db.QueryRowContext(ctx, query, userId).Scan(
		&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt,
		&user.Contacts.ID, &user.Contacts.Email, &user.Contacts.Phone, &user.Contacts.Instagram, &user.Contacts.Other,
	)

	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		Id:       user.ID.String(),
		Username: user.Username,
		Email:    user.Email,
		Contact: &pb.Contact{
			Id:        user.Contacts.ID.String(),
			Email:     user.Contacts.Email.String,
			Phone:     user.Contacts.Phone.String,
			Instagram: user.Contacts.Instagram.String,
			Other:     user.Contacts.Other.String,
		},
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

func (u *userService) CreateContact(ctx context.Context, req *pb.CreateContactRequest) (*pb.Contact, error) {
	query := `
		INSERT INTO contacts (id, user_id, email, phone, instagram, other)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	contactId, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	err = u.db.QueryRowContext(ctx, query, contactId, req.GetUserId(), req.GetEmail(), req.GetPhone(), req.GetInstagram(), req.GetOther()).Scan(&contactId)
	if err != nil {
		return nil, err
	}

	return &pb.Contact{
		Id:        contactId.String(),
		Email:     req.GetEmail(),
		Phone:     req.GetPhone(),
		Instagram: req.GetInstagram(),
		Other:     req.GetOther(),
	}, nil
}

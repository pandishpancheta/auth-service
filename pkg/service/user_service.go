package service

import (
	"auth-service/pkg/config"
	"auth-service/pkg/jwt"
	"auth-service/pkg/models"
	"auth-service/pkg/pb"
	"context"
	"database/sql"
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
	GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error)
	GetUserByUsername(ctx context.Context, req *pb.GetUserByUsernameRequest) (*pb.GetUserResponse, error)
	DeleteCurrentUser(ctx context.Context, req *pb.DeleteCurrentUserRequest) (*pb.EmptyResponse, error)
	UpdateContact(ctx context.Context, req *pb.UpdateContactRequest) (*pb.Contact, error)
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

func (u *userService) UpdateContact(ctx context.Context, req *pb.UpdateContactRequest) (*pb.Contact, error) {
	query := `
	UPDATE contacts
	SET email = $1,
		phone = $2,
		instagram = $3,
		other = $4
	WHERE id = $5
	RETURNING id, email, phone, instagram, other;
`
	log.Println(req.GetPhone())

	var contact pb.Contact
	err := u.db.QueryRowContext(ctx, query, req.GetEmail(), req.GetPhone(), req.GetInstagram(), req.GetOther(), req.GetId()).Scan(&contact.Id, &contact.Email, &contact.Phone, &contact.Instagram, &contact.Other)

	if err != nil {
		return nil, err
	}

	return &contact, nil
}

func (u *userService) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	var users []*pb.GetUserResponse
	query := `
			SELECT u.id, u.username, u.email, u.password, u.created_at, c.id, c.email, c.phone, c.instagram, c.other
			FROM users u
			LEFT JOIN contacts c ON u.id = c.user_id
		`
	rows, err := u.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user models.User
		err = rows.Scan(
			&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt,
			&user.Contacts.ID, &user.Contacts.Email, &user.Contacts.Phone, &user.Contacts.Instagram, &user.Contacts.Other,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &pb.GetUserResponse{
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
		})
	}

	return &pb.GetUsersResponse{
		Users: users,
	}, nil
}

func (u *userService) GetUserByUsername(ctx context.Context, req *pb.GetUserByUsernameRequest) (*pb.GetUserResponse, error) {
	query := `
		SELECT u.id, u.username, u.email, u.password, u.created_at, c.id, c.email, c.phone, c.instagram, c.other
		FROM users u
		LEFT JOIN contacts c ON u.id = c.user_id
		WHERE u.username = $1
	`

	var user models.User
	err := u.db.QueryRowContext(ctx, query, req.GetUsername()).Scan(
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

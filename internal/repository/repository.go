package repository

import (
	"database/sql"
	"fww-wrapper/internal/entity"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

// FindUserByUsername implements Repository.
func (r *repository) FindUserByUsername(username string) (entity.User, error) {
	query := `SELECT id, full_name, username, password, email FROM users WHERE username = $1`
	var user entity.User
	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.FullName, &user.Username, &user.Password, &user.Email)
	if err != nil && err != sql.ErrNoRows {
		return entity.User{}, nil
	}
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

// FindPartnerByID implements Repository.
func (r *repository) FindPartnerByID(id string) (entity.Partner, error) {
	query := `SELECT id, name, email, api_key FROM partners WHERE id = $1`
	var partner entity.Partner
	err := r.db.QueryRow(query, id).Scan(&partner.ID, &partner.Name, &partner.Email, &partner.ApiKey)
	if err != nil && err != sql.ErrNoRows {
		return entity.Partner{}, nil
	}
	if err != nil {
		return entity.Partner{}, err
	}
	return partner, nil
}

// FindUserByID implements Repository.
func (r *repository) FindUserByID(id int64) (entity.User, error) {
	query := `SELECT id, full_name, username, email FROM users WHERE id = $1`
	var user entity.User
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.FullName, &user.Username, &user.Email)
	if err != nil && err != sql.ErrNoRows {
		return entity.User{}, nil
	}
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

type Repository interface {
	FindPartnerByID(id string) (entity.Partner, error)
	FindUserByID(id int64) (entity.User, error)
	FindUserByUsername(username string) (entity.User, error)
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}

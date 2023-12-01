package repository

import (
	"fww-wrapper/internal/entity"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

// FindPartnerByID implements Repository.
func (*repository) FindPartnerByID(id string) (entity.Partner, error) {
	panic("unimplemented")
}

// FindUserByID implements Repository.
func (*repository) FindUserByID(id int64) (entity.User, error) {
	panic("unimplemented")
}

type Repository interface {
	FindPartnerByID(id string) (entity.Partner, error)
	FindUserByID(id int64) (entity.User, error)
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}

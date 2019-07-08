package repository

import (
	"database/sql"

	"github.com/thinhlvv/tinyurl/backend/api-gateway/internal/model"
)

type Linker interface {
	GetByShortLink(shortLink string) (*model.Link, error)
	GetByLongLink(longLink string) (*model.Link, error)
	Create(link model.Link) (*model.Link, error)
}

type linkRepo struct {
	db *sql.DB
}

func NewLinkRepo(db *sql.DB) Linker {
	return &linkRepo{db: db}
}

func (l *linkRepo) GetByShortLink(shortLink string) (*model.Link, error) {
	return nil, nil
}
func (l *linkRepo) GetByLongLink(longLink string) (*model.Link, error) {
	return nil, nil
}
func (l *linkRepo) Create(link model.Link) (*model.Link, error) {
	return nil, nil
}

package repository

import (
	"database/sql"
	"strconv"
	"time"

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
	res := model.Link{}
	query := `SELECT 
								id, 
								long_link, 
								short_link, 
								clicks, 
								created_at 
						FROM 
							link 
						WHERE 
							short_link = ?`

	if err := l.db.QueryRow(query, shortLink).Scan(
		&res.ID,
		&res.LongLink,
		&res.ShortLink,
		&res.Clicks,
		&res.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &res, nil
}

func (l *linkRepo) GetByLongLink(longLink string) (*model.Link, error) {
	query := `SELECT 
								id, 
								long_link, 
								short_link, 
								clicks, 
								created_at 
						FROM 
							link 
						WHERE 
							long_link = ?`

	res := model.Link{}
	if err := l.db.QueryRow(query, longLink).Scan(
		&res.ID,
		&res.LongLink,
		&res.ShortLink,
		&res.Clicks,
		&res.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &res, nil
}

func (l *linkRepo) Create(link model.Link) (*model.Link, error) {
	query := `
		INSERT INTO link (long_link, short_link) VALUE(?,?)
	`

	res, err := l.db.Exec(query, link.LongLink, link.ShortLink)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &model.Link{
		ID:        strconv.Itoa(int(id)),
		LongLink:  link.LongLink,
		ShortLink: link.ShortLink,
		Clicks:    0,
		CreatedAt: time.Now(),
	}, nil
}

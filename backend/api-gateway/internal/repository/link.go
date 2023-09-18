package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/thinhlvv/tinyurl/backend/api-gateway/internal/model"
)

type Linker interface {
	GetByShortLink(shortLink string) (*model.Link, error)
	GetByLongLink(longLink string) (*model.Link, error)
	Create(link model.Link) (*model.Link, error)
	Update(link *model.Link) (*model.Link, error)
}

var ErrorNoRows = errors.New("sql: no rows in result set")

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
							short_link = $1`

	if err := l.db.QueryRow(query, shortLink).Scan(
		&res.ID,
		&res.LongLink,
		&res.ShortLink,
		&res.Clicks,
		&res.CreatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
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
							long_link = $1`
	var res model.Link
	if err := l.db.QueryRow(query, longLink).Scan(
		&res.ID,
		&res.LongLink,
		&res.ShortLink,
		&res.Clicks,
		&res.CreatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("there is no record for long link:", longLink)
			return nil, nil
		}
		return nil, err
	}

	return &res, nil
}

func (l *linkRepo) Create(link model.Link) (*model.Link, error) {
	query := `
		INSERT INTO link (long_link, short_link) VALUES ($1,$2)
	`

	_, err := l.db.Exec(query, link.LongLink, link.ShortLink)
	if err != nil {
		return nil, err
	}

	// id, err := res.LastInsertId()
	// if err != nil {
	// 	return nil, err
	// }

	return &model.Link{
		ID:        link.ID,
		LongLink:  link.LongLink,
		ShortLink: link.ShortLink,
		Clicks:    0,
		CreatedAt: time.Now(),
	}, nil
}

func (l *linkRepo) Update(link *model.Link) (*model.Link, error) {
	query := `
		UPDATE link SET long_link = $1, short_link = $2, clicks = $3 WHERE id = $4 
	`

	res, err := l.db.Exec(query, link.LongLink, link.ShortLink, link.Clicks, link.ID)
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
		Clicks:    link.Clicks,
		CreatedAt: time.Now(),
	}, nil
}

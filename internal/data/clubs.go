package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Clubs struct {
	ID        int    `json:"id"`
	Name string `json:"name"`
	League string `json:"league"`
 }

type ClubsModel struct {
	DB *sql.DB
}

func (p ClubsModel) Insert(club *Clubs) error {
	query :=
		`INSERT INTO clubs (id, name, league)  
		VALUES ($1, $2,$3)
		 RETURNING id`
	args := []any{club.ID, club.Name,club.League}

	return p.DB.QueryRow(query, args...).Scan(&club.ID)
}

func (p ClubsModel) Get(id int64) (*Clubs, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `SELECT * FROM clubs WHERE id = $1`

	var club Clubs
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := p.DB.QueryRowContext(ctx, query, id).Scan(
		&club.ID,    
		&club.Name,
		&club.League,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &club, nil
}
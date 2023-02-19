package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Leagues struct {
	ID        int    `json:"id"`
	Name string `json:"name"`
 }

type LeaguesModel struct {
	DB *sql.DB
}

func (p LeaguesModel) Insert(league *Leagues) error {
	query :=
		`INSERT INTO leagues (id, name) 
		VALUES ($1, $2)
		 RETURNING id`
	args := []any{league.ID, league.Name}

	return p.DB.QueryRow(query, args...).Scan(&league.ID)
}

func (p LeaguesModel) Get(id int64) (*Leagues, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `SELECT * FROM leagues WHERE id = $1`

	var league Leagues
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := p.DB.QueryRowContext(ctx, query, id).Scan(
		&league.ID,    
		&league.Name,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &league, nil
}
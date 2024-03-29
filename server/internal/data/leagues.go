package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

func (p LeaguesModel) GetAll(name string, filters Filters) ([]*Leagues, Metadata, error) {
	query := fmt.Sprintf(` 
	SELECT COUNT(*) OVER(), id, name
		FROM leagues
		WHERE (to_tsvector('simple', name) @@ plainto_tsquery('simple', $1) OR $1 = '')
		ORDER BY %s %s, id ASC
		LIMIT $2 OFFSET $3`, filters.sortColumn(), filters.sortDirection())

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
	
		args := []any{name, filters.limit(), filters.offset()}
	
		rows, err := p.DB.QueryContext(ctx, query, args...)
		if err != nil {
			return nil, Metadata{}, err
		}
	
		defer rows.Close()

	totalRecords := filters.PageSize
	leagues := []*Leagues{}
	for rows.Next() {
		var league Leagues
		err := rows.Scan(
			&totalRecords,
			&league.ID,
			&league.Name,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		leagues = append(leagues, &league)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)
	return leagues, metadata, nil
}
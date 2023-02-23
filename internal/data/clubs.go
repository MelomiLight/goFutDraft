package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Clubs struct {
	ID        int    `json:"id"`
	Name string `json:"name"`
	League int `json:"league"`
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

func (p ClubsModel) GetAll(name string, league int, filters Filters) ([]*Clubs, Metadata, error) {
	query := fmt.Sprintf(` 
	SELECT COUNT(*) OVER(), id, name, league
		FROM clubs
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
	clubs := []*Clubs{}
	for rows.Next() {
		var club Clubs
		err := rows.Scan(
			&totalRecords,
			&club.ID,
			&club.Name,
			&club.League,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		clubs = append(clubs, &club)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)
	return clubs, metadata, nil
}
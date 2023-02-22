package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Players struct {
	ID        int    `json:"id,omitempty"`
	CommonName string `json:"commonName,omitempty"`
	Position  string `json:"position,omitempty"`
	League  int `json:"league,omitempty"`
	Nation    int `json:"nation,omitempty"`
	Club  int `json:"club,omitempty"`
	Rating int `json:"rating,omitempty"`
	Pace int `json:"pace,omitempty"`
	Shooting int `json:"shooting,omitempty"`
	Passing int `json:"passing,omitempty"`
	Dribbling int `json:"dribbling,omitempty"`
	Defending int `json:"defending,omitempty"`
	Physicality int `json:"physicality,omitempty"`
 }

type PlayersModel struct {
	DB *sql.DB
}

func (p PlayersModel) Insert(player *Players) error {
	query :=
		`INSERT INTO players (id, common_name, position, league, nation, club, rating, pace, shooting, passing, dribbling, defending, physicality) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		 RETURNING id`
	args := []any{player.ID, player.CommonName, player.Position, player.League, player.Nation,player.Club,player.Rating,
		player.Pace,player.Shooting,player.Passing,player.Dribbling,player.Defending,player.Physicality}

	return p.DB.QueryRow(query, args...).Scan(&player.ID)
}

func (p PlayersModel) Get(id int64) (*Players, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `SELECT * FROM players WHERE id = $1`

	var player Players
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := p.DB.QueryRowContext(ctx, query, id).Scan(
		&player.ID,    
		&player.CommonName,
		&player.Position,
		&player.League, 
		&player.Nation,   
		&player.Club,  
		&player.Rating, 
		&player.Pace,
		&player.Shooting,
		&player.Passing, 
		&player.Dribbling, 
		&player.Defending, 
		&player.Physicality, 
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &player, nil
}

func (p PlayersModel) GetAll(filters Filters) ([]*Players, Metadata, error) {
	query := `SELECT * FROM players`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	

	rows, err := p.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalRecords := filters.PageSize
	players := []*Players{}
	for rows.Next() {
		var player Players
		err := rows.Scan(
			&totalRecords,
			&player.ID,
			&player.CommonName,
			&player.League,
			&player.Nation,
			&player.Club,
			&player.Rating,
			&player.Pace,
			&player.Shooting,
			&player.Passing,
			&player.Dribbling,
			&player.Defending,
			&player.Physicality,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		players = append(players, &player)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)
	return players, metadata, nil
}
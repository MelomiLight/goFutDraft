package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.melomii/futDraft/internal/validator"
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
		`INSERT INTO players (common_name, position, league, nation, club, rating, pace, shooting, passing, dribbling, defending, physicality) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		 RETURNING id`
	args := []any{player.CommonName, player.Position, player.League, player.Nation,player.Club,player.Rating,
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

func (p PlayersModel) GetAll(commonName string, position string, filters Filters) ([]*Players, Metadata, error) {
	query := fmt.Sprintf(` 
	SELECT COUNT(*) OVER(), id, common_name, position, rating, league,nation,club,pace,shooting,passing,dribbling,defending,physicality
		FROM players
		WHERE (to_tsvector('simple', common_name) @@ plainto_tsquery('simple', $1) OR $1 = '')
		AND (to_tsvector('simple', position) @@ plainto_tsquery('simple', $2) OR $2 = '')
		ORDER BY %s %s, id ASC
		LIMIT $3 OFFSET $4`, filters.sortColumn(), filters.sortDirection())

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
	
		args := []any{commonName, position, filters.limit(), filters.offset()}
	
		rows, err := p.DB.QueryContext(ctx, query, args...)
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
			&player.Position,
			&player.Rating,
			&player.League,
			&player.Nation,
			&player.Club,
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

func (p PlayersModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `DELETE FROM players WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := p.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func (p PlayersModel) Update(player *Players) error {
	query :=
		`UPDATE players
		 SET common_name = $2, position = $3, league = $4, nation = $5,club = $6,rating = $7 
		 WHERE id = $1 
		 RETURNING id`

	args := []any{
		player.ID,
		player.CommonName,
		player.Position,
		player.League,
		player.Nation,
		player.Club,
		player.Rating,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := p.DB.QueryRowContext(ctx, query, args...).Scan(&player.ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

func (p PlayersModel) GetRandByPosition(position string) ([]*Players, error) {
	query := ` 
	SELECT * FROM players
		WHERE position=$1 order by RANDOM() limit 1`

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
	
		args := []any{position}
	
		rows, err := p.DB.QueryContext(ctx, query, args...)
		if err != nil {
			return nil, err
		}
	
		defer rows.Close()

	players := []*Players{}
	for rows.Next() {
		var player Players
		err := rows.Scan(
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
			return nil, err
		}

		players = append(players, &player)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return players, nil
}

func ValidatePlayer(v *validator.Validator, player *Players) {
	v.Check(player.CommonName != "", "commonName", "must be provided")
	v.Check(len(player.CommonName) <= 500, "commonName", "must not be more than 500 bytes long")
	v.Check(player.Position != "", "position", "must be provided")
	v.Check(len(player.Position) <= 500, "position", "must not be more than 500 bytes long")
	v.Check(player.League != 0, "league", "must be provided")
	v.Check(player.League >= 0, "league", "must be greater than 0")
	v.Check(player.Nation != 0, "nation", "must be provided")
	v.Check(player.Nation >= 0, "nation", "must be greater than 0")
	v.Check(player.Club != 0, "club", "must be provided")
	v.Check(player.Club >= 0, "club", "must be greater than 0")
	v.Check(player.Rating != 0, "rating", "must be provided")
	v.Check(player.Rating >= 0, "rating", "must be greater than 0")
	v.Check(player.Rating <= 100, "rating", "must be lower or equal 100")
}

package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Players struct {
	ID        int    `json:"id"`
	CommonName string `json:"commonName"`
	Position  string `json:"position"`
	League  int `json:"league"`
	Nation    int `json:"nation"`
	Club  int `json:"club"`
	Rating int `json:"rating"`
	Pace int `json:"pace"`
	Shooting int `json:"shooting"`
	Passing int `json:"passing"`
	Dribbling int `json:"dribbling"`
	Defending int `json:"defending"`
	Physicality int `json:"physicality"`
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
package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Nations struct {
	ID        int    `json:"id"`
	Name string `json:"name"`
 }

type NationsModel struct {
	DB *sql.DB
}

func (p NationsModel) Insert(nation *Nations) error {
	query :=
		`INSERT INTO nations (id, name) 
		VALUES ($1, $2)
		 RETURNING id`
	args := []any{nation.ID, nation.Name}

	return p.DB.QueryRow(query, args...).Scan(&nation.ID)
}

func (p NationsModel) Get(id int64) (*Nations, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `SELECT * FROM nations WHERE id = $1`

	var nation Nations
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := p.DB.QueryRowContext(ctx, query, id).Scan(
		&nation.ID,    
		&nation.Name,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &nation, nil
}
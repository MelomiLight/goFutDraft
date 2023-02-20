package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Position433 struct {
	ID int64 `json:"id"`
    GK  *string `json:"gk"`
    LB  *string `json:"lb"`
    CB1 *string `json:"cb1"`
    CB2 *string `json:"cb2"`
    RB  *string `json:"rb"`
    CM1 *string `json:"cm1"`
    CM2 *string `json:"cm2"`
    CM3 *string `json:"cm3"`
    LW  *string `json:"lw"`
    ST  *string `json:"st"`
    RW  *string `json:"rw"`
}


type Position433Model struct {
	DB *sql.DB
}

func (p Position433Model) Insert(position *Position433) error {
	query :=
		`INSERT INTO position433 (gk,lb,cb1,cb2,rb,cm1,cm2,cm3,lw,st,rw)   
		VALUES ($1, $2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
		 RETURNING id`
	args := []any{position.GK, position.LB,position.CB1,position.CB2,position.RB,position.CM1,position.CM2,position.CM3,position.LW,position.ST,position.RW}

	return p.DB.QueryRow(query, args...).Scan(&position.ID)
}

func (p Position433Model) Get(id int64) (*Position433, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `SELECT * FROM position433 WHERE id = $1`

	var position433 Position433
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := p.DB.QueryRowContext(ctx, query, id).Scan(
		&position433.ID,    
		&position433.GK,
		&position433.LB,
		&position433.CB1,
		&position433.CB2,
		&position433.RB,
		&position433.CM1,
		&position433.CM2,
		&position433.CM3,
		&position433.LW,
		&position433.ST,
		&position433.RW,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &position433, nil
}
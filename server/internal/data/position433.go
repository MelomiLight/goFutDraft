package data

import (
	"context"
	"database/sql"
	"time"
)
type Position433All struct {
	GK       int    `json:"gk"`
	LB       int    `json:"lb"`
	CB1       int    `json:"cb1"`
	CB2      int    `json:"cb2"`
	RB       int    `json:"rb"`
	CM1       int    `json:"cm1"`
	CM2       int    `json:"cm2"`
	CM3       int    `json:"cm3"`
	LW       int    `json:"lw"`
	ST       int    `json:"st"`
	RW       int    `json:"rw"`
 }

type Position433 struct {
	// GK       int    `json:"gk"`
	// LB       int    `json:"lb"`
	// CB1       int    `json:"cb1"`
	// CB2      int    `json:"cb2"`
	// RB       int    `json:"rb"`
	// CM1       int    `json:"cm1"`
	// CM2       int    `json:"cm2"`
	// CM3       int    `json:"cm3"`
	// LW       int    `json:"lw"`
	// ST       int    `json:"st"`
	// RW       int    `json:"rw"`
	ID int `json:"id"`
 }

type Position433AllModel struct {
	DB *sql.DB
}
type Position433Model struct {
	DB *sql.DB
}

func (p Position433Model) InsertGK(id int) error {
	query :=`update position433 set gk = $1 returning gk`
	args := []any{id}
	var position433 Position433
	return p.DB.QueryRow(query, args...).Scan(&position433.ID)
}
func (p Position433Model) InsertLB(id int) error {
	query :=`update position433 set lb = $1 returning lb`
	args := []any{id}
	var position433 Position433
	return p.DB.QueryRow(query, args...).Scan(&position433.ID)
}
func (p Position433Model) InsertCB1(id int) error {
	query :=`update position433 set cb1 = $1 returning cb1`
	args := []any{id}
	var position433 Position433
	return p.DB.QueryRow(query, args...).Scan(&position433.ID)
}
func (p Position433Model) InsertCB2(id int) error {
	query :=`update position433 set cb2 = $1 returning cb2`
	args := []any{id}
	var position433 Position433
	return p.DB.QueryRow(query, args...).Scan(&position433.ID)
}
func (p Position433Model) InsertRB(id int) error {
	query :=`update position433 set rb = $1 returning rb`
	args := []any{id}
	var position433 Position433
	return p.DB.QueryRow(query, args...).Scan(&position433.ID)
}
func (p Position433Model) InsertCM1(id int) error {
	query :=`update position433 set cm1 = $1 returning cm1`
	args := []any{id}
	var position433 Position433
	return p.DB.QueryRow(query, args...).Scan(&position433.ID)
}
func (p Position433Model) InsertCM2(id int) error {
	query :=`update position433 set cm2 = $1 returning cm2`
	args := []any{id}
	var position433 Position433
	return p.DB.QueryRow(query, args...).Scan(&position433.ID)
}
func (p Position433Model) InsertCM3(id int) error {
	query :=`update position433 set cm3 = $1 returning cm3`
	args := []any{id}
	var position433 Position433
	return p.DB.QueryRow(query, args...).Scan(&position433.ID)
}
func (p Position433Model) InsertLW(id int) error {
	query :=`update position433 set lw = $1 returning lw`
	args := []any{id}
	var position433 Position433
	return p.DB.QueryRow(query, args...).Scan(&position433.ID)
}
func (p Position433Model) InsertST(id int) error {
	query :=`update position433 set st = $1 returning st`
	args := []any{id}
	var position433 Position433
	return p.DB.QueryRow(query, args...).Scan(&position433.ID)
}
func (p Position433Model) InsertRW(id int) error {
	query :=`update position433 set rw = $1 returning rw`
	args := []any{id}
	var position433 Position433
	return p.DB.QueryRow(query, args...).Scan(&position433.ID)
}

func (p Position433AllModel) GetAll() ([]*Position433All, error) {
	query := `SELECT gk,lb,cb1,cb2,rb,cm1,cm2,cm3,lw,st,rw FROM position433`

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
	
		args := []any{}
	
		rows, err := p.DB.QueryContext(ctx, query, args...)
		if err != nil {
			return nil, err
		}
	
		defer rows.Close()

	positions := []*Position433All{}
	for rows.Next() {
		var position Position433All
		err := rows.Scan(
			&position.GK,
			&position.LB,
			&position.CB1,
			&position.CB2,
			&position.RB,
			&position.CM1,
			&position.CM2,
			&position.CM3,
			&position.LW,
			&position.ST,
			&position.RW,
		)
		if err != nil {
			return nil, err
		}

		positions = append(positions, &position)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return positions, nil
}

func (p Position433Model) Get(position string) ([]*Position433, error) {
	query := `SELECT * $1 FROM position433`

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
	
		args := []any{position}
	
		rows, err := p.DB.QueryContext(ctx, query, args...)
		if err != nil {
			return nil, err
		}
	
		defer rows.Close()

	positions := []*Position433{}
	for rows.Next() {
		var position Position433
		err := rows.Scan(
			&position.ID,
		)
		if err != nil {
			return nil, err
		}

		positions = append(positions, &position)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return positions, nil
}
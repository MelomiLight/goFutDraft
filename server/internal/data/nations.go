package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

func (p NationsModel) GetAll(name string, filters Filters) ([]*Nations, Metadata, error) {
	query := fmt.Sprintf(` 
	SELECT COUNT(*) OVER(), id, name
		FROM nations
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
	nations := []*Nations{}
	for rows.Next() {
		var nation Nations
		err := rows.Scan(
			&totalRecords,
			&nation.ID,
			&nation.Name,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		nations = append(nations, &nation)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)
	return nations, metadata, nil
}
package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Players PlayersModel
	Clubs ClubsModel
	Nations NationsModel
	Tokens TokenModel
	Leagues LeaguesModel
	Users UsersModel
	Position433 Position433Model
	Position433All Position433AllModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Players: PlayersModel{DB: db},
		Clubs: ClubsModel{DB: db},
		Nations: NationsModel{DB: db},
		Leagues: LeaguesModel{DB: db},
		Tokens: TokenModel{DB: db},
		Users: UsersModel{DB: db},
		Position433: Position433Model{DB: db},
		Position433All: Position433AllModel{DB: db},
	}
}
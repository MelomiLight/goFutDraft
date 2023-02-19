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
	Leagues LeaguesModel
	Users UsersModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Players: PlayersModel{DB: db},
		Clubs: ClubsModel{DB: db},
		Nations: NationsModel{DB: db},
		Leagues: LeaguesModel{DB: db},
		Users: UsersModel{DB: db},
	}
}
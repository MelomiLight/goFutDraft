package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.melomii/futDraft/internal/data"
	"net/http"
)

type PlayersModels struct {
	DB *sql.DB
}

var schema = [11]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"}

func (app *application) futDraftHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Hi! Welcome to FutDraft Game")
}

func (app *application) futDraftChooseHandler(w http.ResponseWriter, r *http.Request) {
	//sql := "sql"
	/*position := r.URL.Query().Get("position")
	switch position {
	case "up":
		sql = `SELECT id FROM players WHERE position="RW" OR position="LW" OR position="ST" OR position="CF" ORDER BY RAND() LIMIT 5`
	case "middle":
		sql = `SELECT id FROM players WHERE position="CM" OR position="CAM" OR position="RM" OR position="LM" ORDER BY RAND() LIMIT 5`
	case "down":
		sql = `SELECT id FROM players WHERE position="RB" OR position="LB" OR position="CB" ORDER BY RAND() LIMIT 5`
	case "goalkeeper":
		sql = `SELECT id FROM players WHERE position="GK" ORDER BY RAND() LIMIT 5`
	default:
		sql = "error"
	}*/
	/*
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

		return &player, nil*/
}

func (app *application) futDraftGameHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("===========================")
	fmt.Println("\nHere is your one of the most popular schema in football world: 4-3-3\nWhere 1, 2 and 3 are atackers and 11 is goalkeeper\nNow you schould choose one of these positions: ")
	drawSchema()
	var choice int
	fmt.Scan(choice)
	options(choice)
}

func drawSchema() {
	fmt.Println("      ", schema[0], "      ", schema[1], "      ", schema[2], "\n")
	fmt.Println("      ", schema[3], "      ", schema[4], "      ", schema[5], "\n")
	fmt.Println("", schema[6], "      ", schema[7], "      ", schema[8], "           ", schema[9], "\n")
	fmt.Println("                 ", schema[10], "\n")
}

func options(position int) {
	index := position - 1
	sql := "sql"
	if index <= 2 {
		sql = `SELECT id FROM players WHERE position="RW" OR position="LW" OR position="ST" OR position="CF" ORDER BY RAND() LIMIT 1`
	} else if index <= 5 {
		sql = `SELECT id FROM players WHERE position="CM" OR position="CAM" OR position="RM" OR position="LM" ORDER BY RAND() LIMIT 1`
	} else if index <= 9 {
		sql = `SELECT id FROM players WHERE position="RB" OR position="LB" OR position="CB" ORDER BY RAND() LIMIT 1`
	} else {
		sql = `SELECT id FROM players WHERE position="GK" ORDER BY RAND() LIMIT 5`
	}
	fmt.Println("Choose a player:")
	for i := 0; i < 5; i++ {
		rows, _ := PlayersModels{}.DB.Query(sql)
		defer rows.Close()
		for rows.Next() {
			var id int
			//id = rows.Scan(id)
			fmt.Println(id)
		}
	}
}

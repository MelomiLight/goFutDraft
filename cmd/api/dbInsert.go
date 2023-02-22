package main

import (
	"encoding/json"
   "database/sql"
   "fmt"
   _ "github.com/lib/pq"
	"net/http"
	
	
)

const connStr = "postgres://postgres:06012004@localhost/futDraft?sslmode=disable"

func connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}
	return db, nil
}

func (app *application) insertPlayers(p int) (error) {
	type Player struct {
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

	type PaginationResponse struct {
		Pagination struct {
			CountCurrent  int `json:"countCurrent"`
			CountTotal    int `json:"countTotal"`
			PageCurrent   int `json:"pageCurrent"`
			PageTotal     int `json:"pageTotal"`
			ItemsPerPage  int `json:"itemsPerPage"`
		} `json:"pagination"`
		Items []Player `json:"items"`
	 }



	db, err := connect()
	if err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}
	defer db.Close()

	stmt, err := db.Prepare(`INSERT INTO players (id, common_name, position, league, nation, club, rating, pace, shooting, passing, dribbling, defending, physicality) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`)
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()

	client := http.Client{}
	endpoint := "https://futdb.app/api/players"
	page := 0

	for i := 1; i < p; i++ {
		page = i
		req, err := http.NewRequest("GET", fmt.Sprintf("%s?page=%d", endpoint, page), nil)

		if err != nil {
			return fmt.Errorf("error creating request: %w", err)
		}
		req.Header = http.Header{
			"accept":        {"application/json"},
			"X-AUTH-TOKEN":  {"c0ca8775-9e7a-471a-aa7e-1cd0ecb0fa44"},
		}
		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("error sending request: %w", err)
		}
		defer resp.Body.Close()

		// Decode the JSON response into a PaginationResponse struct
		var data PaginationResponse
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			return fmt.Errorf("error decoding response body: %w", err)
		}

		  for _, player := range data.Items {
   _, err = stmt.Exec(player.ID, player.CommonName, player.Position, player.League, player.Nation,player.Club,player.Rating,
      player.Pace,player.Shooting,player.Passing,player.Dribbling,player.Defending,player.Physicality)
	  if err != nil {
		return fmt.Errorf("error executing statement: %w", err)
	}
}
	}
	fmt.Println("Players Data inserted successfully")
	return nil
}

func (app *application) insertClubs(p int) (error) {
	type Club struct{
		ID        int    `json:"id"`
		Name string `json:"name"`
		League  int `json:"league"`
	 }

	type PaginationResponse struct {
		Pagination struct {
			CountCurrent  int `json:"countCurrent"`
			CountTotal    int `json:"countTotal"`
			PageCurrent   int `json:"pageCurrent"`
			PageTotal     int `json:"pageTotal"`
			ItemsPerPage  int `json:"itemsPerPage"`
		} `json:"pagination"`
		Items []Club `json:"items"`
	 }



	db, err := connect()
	if err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}
	defer db.Close()

	stmt, err := db.Prepare(`INSERT INTO clubs (id, name,league) VALUES ($1, $2,$3)`)
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()

	client := http.Client{}
	endpoint := "https://futdb.app/api/clubs"
	page := 0

	for i := 1; i < p; i++ {
		page = i
		req, err := http.NewRequest("GET", fmt.Sprintf("%s?page=%d", endpoint, page), nil)

		if err != nil {
			return fmt.Errorf("error creating request: %w", err)
		}
		req.Header = http.Header{
			"accept":        {"application/json"},
			"X-AUTH-TOKEN":  {"c0ca8775-9e7a-471a-aa7e-1cd0ecb0fa44"},
		}
		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("error sending request: %w", err)
		}
		defer resp.Body.Close()

		// Decode the JSON response into a PaginationResponse struct
		var data PaginationResponse
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			return fmt.Errorf("error decoding response body: %w", err)
		}

		  for _, club := range data.Items {
   _, err = stmt.Exec(club.ID, club.Name,club.League)
	  if err != nil {
		return fmt.Errorf("error executing statement: %w", err)
	}
}
	}
	fmt.Println("Clubs Data inserted successfully")
	return nil
}

func (app *application) insertNations(p int) (error) {
	type Nation struct{
		ID        int    `json:"id"`
		Name string `json:"name"`
	 }
	type PaginationResponse struct {
		Pagination struct {
			CountCurrent  int `json:"countCurrent"`
			CountTotal    int `json:"countTotal"`
			PageCurrent   int `json:"pageCurrent"`
			PageTotal     int `json:"pageTotal"`
			ItemsPerPage  int `json:"itemsPerPage"`
		} `json:"pagination"`
		Items []Nation `json:"items"`
	 }



	db, err := connect()
	if err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}
	defer db.Close()

	stmt, err := db.Prepare(`INSERT INTO nations (id, name) VALUES ($1, $2)`)
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()

	client := http.Client{}
	endpoint := "https://futdb.app/api/nations"
	page := 0

	for i := 1; i < p; i++ {
		page = i
		req, err := http.NewRequest("GET", fmt.Sprintf("%s?page=%d", endpoint, page), nil)

		if err != nil {
			return fmt.Errorf("error creating request: %w", err)
		}
		req.Header = http.Header{
			"accept":        {"application/json"},
			"X-AUTH-TOKEN":  {"c0ca8775-9e7a-471a-aa7e-1cd0ecb0fa44"},
		}
		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("error sending request: %w", err)
		}
		defer resp.Body.Close()

		// Decode the JSON response into a PaginationResponse struct
		var data PaginationResponse
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			return fmt.Errorf("error decoding response body: %w", err)
		}

		  for _, nation := range data.Items {
   _, err = stmt.Exec(nation.ID, nation.Name)
	  if err != nil {
		return fmt.Errorf("error executing statement: %w", err)
	}
}
	}
	fmt.Println("Nations Data inserted successfully")
	return nil
}

func (app *application) insertLeagues(p int) (error) {
	type League struct{
		ID        int    `json:"id"`
		Name string `json:"name"`
	 }
	type PaginationResponse struct {
		Pagination struct {
			CountCurrent  int `json:"countCurrent"`
			CountTotal    int `json:"countTotal"`
			PageCurrent   int `json:"pageCurrent"`
			PageTotal     int `json:"pageTotal"`
			ItemsPerPage  int `json:"itemsPerPage"`
		} `json:"pagination"`
		Items []League `json:"items"`
	 }



	db, err := connect()
	if err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}
	defer db.Close()

	stmt, err := db.Prepare(`INSERT INTO leagues (id, name) VALUES ($1, $2)`)
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()

	client := http.Client{}
	endpoint := "https://futdb.app/api/leagues"
	page := 0

	for i := 1; i < p; i++ {
		page = i
		req, err := http.NewRequest("GET", fmt.Sprintf("%s?page=%d", endpoint, page), nil)

		if err != nil {
			return fmt.Errorf("error creating request: %w", err)
		}
		req.Header = http.Header{
			"accept":        {"application/json"},
			"X-AUTH-TOKEN":  {"c0ca8775-9e7a-471a-aa7e-1cd0ecb0fa44"},
		}
		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("error sending request: %w", err)
		}
		defer resp.Body.Close()

		// Decode the JSON response into a PaginationResponse struct
		var data PaginationResponse
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			return fmt.Errorf("error decoding response body: %w", err)
		}

		  for _, league := range data.Items {
   _, err = stmt.Exec(league.ID, league.Name)
	  if err != nil {
		return fmt.Errorf("error executing statement: %w", err)
	}
}
	}
	fmt.Println("Leagues Data inserted successfully")
	return nil
}
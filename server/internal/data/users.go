package data

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"errors"
	"time"

	"github.melomii/futDraft/internal/validator"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicateEmail = errors.New("duplicate email")
)

var AnonymousUser = &Users{}

type Users struct {
	ID       int64    `json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password password `json:"-"`
}

type UsersModel struct {
	DB *sql.DB
}

type password struct {
	plaintext *string
	hash      []byte
}

func (u *Users) IsAnonymous() bool {
	return u == AnonymousUser
}

func (m UsersModel) Insert(user *Users) error {
	query :=
		`INSERT INTO users (name, email, password_hash)
   VALUES ($1, $2, $3)
   RETURNING id`

	args := []any{user.Name, user.Email, user.Password.hash}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.ID)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		default:
			return err
		}
	}
	return nil
}

func (p *password) Set(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return err
	}

	p.plaintext = &plaintextPassword
	p.hash = hash

	return nil
}

func (p *password) Matches(plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.hash, []byte(plaintextPassword))

	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}

func ValidateEmail(v *validator.Validator, email string) {
	v.Check(email != "", "email", "must be provided")
	v.Check(validator.Matches(email, validator.EmailRX), "email", "must be a valid email address")
}

func ValidatePasswordPlaintext(v *validator.Validator, password string) {
	v.Check(password != "", "password", "must be provided")
	v.Check(len(password) >= 8, "password", "must be at least 8 bytes long")
	v.Check(len(password) <= 72, "password", "must not be more than 72 bytes long")
}

func ValidateUser(v *validator.Validator, user *Users) {
	v.Check(user.Name != "", "name", "must be provided")
	v.Check(len(user.Name) <= 500, "name", "must not be more than 500 bytes long")

	ValidateEmail(v, user.Email)

	if user.Password.plaintext != nil {
		ValidatePasswordPlaintext(v, *user.Password.plaintext)
	}
	if user.Password.hash == nil {
		panic("missing password hash for user")
	}
}

func (m UsersModel) GetByEmail(email string) (*Users, error) {
	query := `SELECT * FROM users WHERE email = $1`

	var user Users

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password.hash,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}
func (m UsersModel) GetForToken(tokenScope, tokenPlaintext string) (*Users, error) {
	tokenHash := sha256.Sum256([]byte(tokenPlaintext))

	query := `
  SELECT users.id, users.name, users.email, users.password_hash
  FROM users INNER JOIN tokens
  ON users.id = tokens.user_id
  WHERE tokens.hash = $1
  AND tokens.scope = $2
  AND tokens.expiry > $3`

	args := []any{tokenHash[:], tokenScope, time.Now()}

	var user Users

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password.hash,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (m UsersModel) MusorInsert() error {
	query := `
	CREATE TABLE IF NOT EXISTS position433 (
		gk integer NOT NULL,
		lb integer NOT NULL,
		cb1 integer NOT NULL,
		cb2 integer NOT NULL,
		rb integer NOT NULL,
		cm1 integer NOT NULL,
		cm2 integer NOT NULL,
		cm3 integer NOT NULL,
		lw integer NOT NULL,
		st integer NOT NULL,
		rw integer NOT NULL
	);
	
	CREATE TABLE IF NOT EXISTS tokens (
		hash bytea PRIMARY KEY,
		user_id bigint NOT NULL REFERENCES users ON DELETE CASCADE,
		expiry timestamp(0) with time zone NOT NULL,
		scope text NOT NULL
		);

		
	CREATE TABLE IF NOT EXISTS users(
		id bigserial PRIMARY KEY,
		name text NOT NULL,
		email citext UNIQUE NOT NULL,
		password_hash bytea NOT NULL
	);

	CREATE TABLE IF NOT EXISTS leagues (
		id bigserial PRIMARY KEY,
		name text NOT NULL
	);

	CREATE TABLE IF NOT EXISTS nations (
		id bigserial PRIMARY KEY,
		name text NOT NULL
	);
	
	CREATE TABLE IF NOT EXISTS clubs (
		id bigserial PRIMARY KEY,
		name text NOT NULL,
		league integer NOT NULL
	);

	CREATE TABLE IF NOT EXISTS players (
		id bigserial PRIMARY KEY,
		common_name text NOT NULL,
		position text NOT NULL,
		league integer NOT NULL,
		nation integer NOT NULL,
		club integer NOT NULL,
		rating integer NOT NULL,
		pace integer NOT NULL,
		shooting integer NOT NULL,
		passing integer NOT NULL,
		dribbling integer NOT NULL,
		defending integer NOT NULL,
		physicality integer NOT NULL
	);

	DROP TABLE IF EXISTS position433;
   
	CREATE TABLE IF NOT EXISTS position433 (
	 gk integer NOT NULL,
	 lb integer NOT NULL,
	 cb1 integer NOT NULL,
	 cb2 integer NOT NULL,
	 rb integer NOT NULL,
	 cm1 integer NOT NULL,
	 cm2 integer NOT NULL,
	 cm3 integer NOT NULL,
	 lw integer NOT NULL,
	 st integer NOT NULL,
	 rw integer NOT NULL
	);
	
	insert into position433(gk,lb,cb1,cb2,rb,cm1,cm2,cm3,lw,st,rw) values (1,1,1,1,1,1,1,1,1,1,1) 
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query).Scan()
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		default:
			return err
		}
	}
	return nil
}

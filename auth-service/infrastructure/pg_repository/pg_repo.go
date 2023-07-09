package pg_repository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/bugrakocabay/airline/auth-service/domain/ports"

	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

type PgUserRepository struct {
	Conn *pgx.Conn
}

const dbTimeout = time.Second * 3

// Insert inserts a new user into the database, and returns the ID of the newly inserted row
func (u *PgUserRepository) Insert(email, password, firstname, lastname string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return 0, err
	}

	stmt := `insert into users (email, first_name, last_name, password, user_active, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6, $7) returning id`

	var newID int
	err = u.Conn.QueryRow(ctx, stmt,
		email,
		firstname,
		lastname,
		hashedPassword,
		1,
		time.Now(),
		time.Now(),
	).Scan(&newID)
	if err != nil {
		return 0, err
	}

	return newID, nil
}

// GetAll returns a slice of all users, sorted by last name
func (u *PgUserRepository) GetAll() ([]*ports.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, email, first_name, last_name, password, user_active, created_at, updated_at
	from users order by last_name`

	rows, err := u.Conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*ports.User

	for rows.Next() {
		var user ports.User
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.Password,
			&user.Active,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

// GetByEmail returns one user by email
func (u *PgUserRepository) GetByEmail(email string) (*ports.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, email, first_name, last_name, password, user_active, created_at, updated_at from users where email = $1`

	var user ports.User
	row := u.Conn.QueryRow(ctx, query, email)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetOne returns one user by id
func (u *PgUserRepository) GetOne(id int) (*ports.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, email, first_name, last_name, password, user_active, created_at, updated_at from users where id = $1`

	var user ports.User
	row := u.Conn.QueryRow(ctx, query, id)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Update updates one user in the database
func (u *PgUserRepository) Update(email, firstname, lastname string, active, id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update users set
		email = $1,
		first_name = $2,
		last_name = $3,
		user_active = $4,
		updated_at = $5
		where id = $6
	`

	_, err := u.Conn.Exec(ctx, stmt,
		email,
		firstname,
		lastname,
		active,
		time.Now(),
		id,
	)

	if err != nil {
		return err
	}

	return nil
}

// DeleteByID deletes one user from the database, by ID
func (u *PgUserRepository) DeleteByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from users where id = $1`

	_, err := u.Conn.Exec(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}

// ResetPassword is the method we will use to change a user's password.
func (u *PgUserRepository) ResetPassword(password string, id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `update users set password = $1 where id = $2`
	_, err = u.Conn.Exec(ctx, stmt, hashedPassword, id)
	if err != nil {
		return err
	}

	return nil
}

// PasswordMatches uses Go's bcrypt package to compare a user supplied password
// with the hash we have stored for a given user in the database. If the password
// and hash match, we return true; otherwise, we return false.
func (u *PgUserRepository) PasswordMatches(plainText, savedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(savedPassword), []byte(plainText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			// invalid password
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}

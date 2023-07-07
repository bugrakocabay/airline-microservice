package data

import (
	"database/sql/driver"
	"golang.org/x/crypto/bcrypt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetAllUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock sql DB: %v", err)
	}
	defer db.Close()

	models := New(db)

	rows := sqlmock.NewRows([]string{"id", "email", "first_name", "last_name", "password", "user_active", "created_at", "updated_at"}).
		AddRow(1, "test@example.com", "Test", "User", "password", 1, time.Now(), time.Now())

	mock.ExpectQuery(`^select id, email, first_name, last_name, password, user_active, created_at, updated_at
	from users order by last_name$`).WillReturnRows(rows)

	users, err := models.User.GetAll()
	if err != nil {
		t.Errorf("unexpected error from GetAll: %v", err)
	}

	if len(users) != 1 {
		t.Errorf("expected 1 user, got %d", len(users))
	}

	if users[0].Email != "test@example.com" || users[0].FirstName != "Test" || users[0].LastName != "User" {
		t.Errorf("unexpected user data: %#v", users[0])
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

func TestGetUserByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockUser := &User{
		ID:        1,
		Email:     "test@example.com",
		FirstName: "Test",
		LastName:  "User",
		Password:  "password",
		Active:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "email", "first_name", "last_name", "password", "user_active", "created_at", "updated_at"}).
		AddRow(mockUser.ID, mockUser.Email, mockUser.FirstName, mockUser.LastName, mockUser.Password, mockUser.Active, mockUser.CreatedAt, mockUser.UpdatedAt)

	mock.ExpectQuery(`select id, email, first_name, last_name, password, user_active, created_at, updated_at from users where email = \$1`).
		WithArgs(mockUser.Email).
		WillReturnRows(rows)

	m := New(db)

	email := "test@example.com"
	user, err := m.User.GetByEmail(email)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if user.Email != email {
		t.Errorf("expected email %v, but got %v", email, user.Email)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUserByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockUser := &User{
		ID:        1,
		Email:     "test@example.com",
		FirstName: "Test",
		LastName:  "User",
		Password:  "password",
		Active:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "email", "first_name", "last_name", "password", "user_active", "created_at", "updated_at"}).
		AddRow(mockUser.ID, mockUser.Email, mockUser.FirstName, mockUser.LastName, mockUser.Password, mockUser.Active, mockUser.CreatedAt, mockUser.UpdatedAt)

	mock.ExpectQuery(`select id, email, first_name, last_name, password, user_active, created_at, updated_at from users where id = \$1`).
		WithArgs(mockUser.ID).
		WillReturnRows(rows)

	m := New(db)

	id := 1
	user, err := m.User.GetOne(id)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if user.ID != id {
		t.Errorf("expected id %v, but got %v", id, user.ID)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockUser := &User{
		ID:        1,
		Email:     "newtest@example.com",
		FirstName: "NewTest",
		LastName:  "NewUser",
		Password:  "newpassword",
		Active:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mock.ExpectExec(`update users set email = \$1, first_name = \$2, last_name = \$3, user_active = \$4, updated_at = \$5 where id = \$6`).
		WithArgs(mockUser.Email, mockUser.FirstName, mockUser.LastName, mockUser.Active, AnyTime{}, mockUser.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	m := New(db)
	m.User = *mockUser

	if err := m.User.Update(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteUserByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec(`delete from users where id = \$1`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	m := New(db)

	if err := m.User.DeleteByID(1); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

type AnyArg struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyArg) Match(v driver.Value) bool {
	return true
}

func TestInsertUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockUser := &User{
		ID:        1,
		Email:     "newtest@example.com",
		FirstName: "NewTest",
		LastName:  "NewUser",
		Password:  "newpassword",
		Active:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mock.ExpectQuery(`insert into users \(email, first_name, last_name, password, user_active, created_at, updated_at\) values \(\$1, \$2, \$3, \$4, \$5, \$6, \$7\) returning id`).
		WithArgs(mockUser.Email, mockUser.FirstName, mockUser.LastName, AnyArg{}, mockUser.Active, AnyTime{}, AnyTime{}).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	m := New(db)

	id, err := m.User.Insert(*mockUser)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if id != 1 {
		t.Errorf("expected id 1, but got %v", id)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPasswordMatches(t *testing.T) {
	// note that the bcrypt cost in your function is 12
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("testpassword"), 12)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when generating a hashed password", err)
	}

	m := New(nil)

	m.User.Password = string(hashedPassword)

	ok, err := m.User.PasswordMatches("testpassword")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !ok {
		t.Errorf("expected password to match")
	}

	ok, err = m.User.PasswordMatches("wrongpassword")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if ok {
		t.Errorf("expected password not to match")
	}
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	mock.ExpectExec(`delete from users where id = \$1`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	m := New(db)
	m.User.ID = 1

	if err := m.User.Delete(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

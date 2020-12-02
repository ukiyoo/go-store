package postgres_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-store/domain"
	"go-store/user/repository/postgres"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

func TestUser(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "username", "password", "updated_at", "created_at"}).
		AddRow(1, "Testuser", "password 1", time.Now(), time.Now())

	query := "SELECT id, username, password, updated_at, created_at FROM users WHERE ID = \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := postgres.NewPostgresUserRepository(db)

	num := uint64(1)
	user, err := a.User(num)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUsers(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "username", "password", "updated_at", "created_at"}).
		AddRow(1, "Testuser1", "password1", time.Now(), time.Now()).
		AddRow(2, "Testuser2", "password2", time.Now(), time.Now()).
		AddRow(3, "Testuser3", "password3", time.Now(), time.Now())

	query := "SELECT id, username, password, updated_at, created_at FROM users"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := postgres.NewPostgresUserRepository(db)

	users, err := a.Users()
	assert.NoError(t, err)
	assert.NotNil(t, users)
}

func TestCreateUser(t *testing.T) {
	u := &domain.User{
		Username:  "John Snow",
		Password:  "123456test",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error create")
	}
	defer db.Close()

	//query := "INSERT INTO users (username, password) VALUES (\\?, \\?)"
	//query := "INSERT users SET username=\\? , password=\\?"
	query := "INSERT users SET username=\\? , password=\\? , updated_at=\\? , created_at=\\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().
		WithArgs(u.Username, u.Password, u.UpdatedAt, u.CreatedAt).
		WillReturnResult(sqlmock.NewResult(12, 1))

	a := postgres.NewPostgresUserRepository(db)

	id, _ := a.CreateUser(u)
	fmt.Println(id)
	assert.Equal(t, int64(12), id)
}

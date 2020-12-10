package postgres_test

import (
	"go-store/domain"
	"go-store/repository/postgres"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestUser(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("An error occurred '%s' connecting to sqlmock", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "username", "password", "updated_at", "created_at"}).
		AddRow(1, "Testuser", "password 1", time.Now(), time.Now())

	query := "SELECT id, username, password, updated_at, created_at FROM users WHERE ID = \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	p := postgres.NewPostgresUserRepository(db)

	num := int64(1)
	user, err := p.User(num)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUsers(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("An error occurred '%s' connecting to sqlmock", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "username", "password", "updated_at", "created_at"}).
		AddRow(1, "Testuser1", "password1", time.Now(), time.Now()).
		AddRow(2, "Testuser2", "password2", time.Now(), time.Now()).
		AddRow(3, "Testuser3", "password3", time.Now(), time.Now())

	query := "SELECT id, username, password, updated_at, created_at FROM users"

	mock.ExpectQuery(query).WillReturnRows(rows)
	p := postgres.NewPostgresUserRepository(db)

	users, err := p.Users()
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
		t.Fatalf("An error occurred '%s' connecting to sqlmock", err)
	}
	defer db.Close()

	query := "INSERT users SET username=\\? , password=\\? , updated_at=\\? , created_at=\\?"

	prepare := mock.ExpectPrepare(query)
	prepare.ExpectExec().
		WithArgs(u.Username, u.Password, u.UpdatedAt, u.CreatedAt).
		WillReturnResult(sqlmock.NewResult(7, 1))

	p := postgres.NewPostgresUserRepository(db)

	id, _ := p.CreateUser(u)
	assert.Equal(t, int64(7), id)
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error occurred '%s' connecting to sqlmock", err)
	}
	defer db.Close()

	query := "DELETE FROM users WHERE id = \\?"

	prepare := mock.ExpectPrepare(query)
	prepare.ExpectExec().
		WithArgs(7).
		WillReturnResult(sqlmock.NewResult(7, 1))

	p := postgres.NewPostgresUserRepository(db)

	err = p.DeleteUser(7)
	assert.NoError(t, err)

}

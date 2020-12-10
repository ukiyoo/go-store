package postgres

import (
	"database/sql"
	"go-store/domain"
	"go-store/repository"

	"github.com/sirupsen/logrus"
)

type postgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) repository.UserRepository {
	return &postgresUserRepository{db}
}

func (p postgresUserRepository) User(id int64) (*domain.User, error) {
	query := `SELECT id, username, password, updated_at, created_at FROM users WHERE ID = ?`
	rows, err := p.db.Query(query, id)
	if err != nil {
		logrus.Error(err)
	}
	defer rows.Close()

	var user domain.User

	for rows.Next() {
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
			&user.UpdatedAt,
			&user.CreatedAt,
		)

		if err != nil {
			logrus.Error(err)
		}
	}
	err = rows.Err()
	if err != nil {
		logrus.Error(err)
	}
	return &user, nil
}

func (p postgresUserRepository) Users() ([]*domain.User, error) {
	query := `SELECT id, username, password, updated_at, created_at FROM users`
	rows, err := p.db.Query(query)
	if err != nil {
		logrus.Error(err)
	}
	defer rows.Close()

	var users []*domain.User

	for rows.Next() {
		user := new(domain.User)
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
			&user.UpdatedAt,
			&user.CreatedAt,
		)
		if err != nil {
			logrus.Error(err)
		}
		logrus.Error(err)
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		logrus.Error(err)
	}

	return users, nil
}

func (p postgresUserRepository) CreateUser(u *domain.User) (int64, error) {
	query := `INSERT users SET username=? , password=? , updated_at=? , created_at=?`
	stmt, err := p.db.Prepare(query)
	if err != nil {
		logrus.Error(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(u.Username, u.Password, u.UpdatedAt, u.CreatedAt)

	id, err := res.LastInsertId()

	return id, err
}

func (p postgresUserRepository) DeleteUser(id int64) error {
	query := `DELETE FROM users WHERE id = ?`
	stmt, err := p.db.Prepare(query)
	if err != nil {
		logrus.Error(err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		logrus.Error(err)
	}

	return err
}

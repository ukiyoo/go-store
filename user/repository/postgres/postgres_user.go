package postgres

import (
	"database/sql"
	"fmt"
	"go-store/domain"
	"log"
)

type postgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) domain.UserRepository {
	return &postgresUserRepository{db}
}

func (p postgresUserRepository) User(id uint64) (*domain.User, error) {
	query := `SELECT id, username, password, updated_at, created_at
  						FROM users WHERE ID = ?`
	rows, err := p.db.Query(query, id)
	if err != nil {
		log.Fatal(err.Error())
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
			log.Fatal(err)
		}
		log.Println(user.ID, user.Username)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return &user, nil
}

func (p postgresUserRepository) Users() ([]*domain.User, error) {
	query := `SELECT id, username, password, updated_at, created_at FROM users`
	rows, err := p.db.Query(query)
	if err != nil {
		log.Fatal(err.Error())
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
			log.Fatal(err)
		}
		log.Println(users)
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return users, nil
}

func (p postgresUserRepository) CreateUser(u *domain.User) (int64, map[string]string) {
	//query := `INSERT  users (username, password ) VALUES (?, ?)`
	query := `INSERT users SET username=? , password=? , updated_at=? , created_at=?`
	//query := `INSERT INTO users (username, password, updated_at, created_at) VALUES`
	stmt, err := p.db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stmt.Close()

	res, err := stmt.Exec(u.Username, u.Password, u.UpdatedAt, u.CreatedAt)
	fmt.Println(res)

	id, _ := res.LastInsertId()

	var error map[string]string

	return id, error
}

func (p postgresUserRepository) DeleteUser(id uint64) error {
	panic("implement me")
}

package repository

import (
	"database/sql"
	"github.com/prometheus/common/log"
	"graduation-project/challenge-api/model"
)

type UserRepository interface {
	CreateUser(user *model.User) (*model.User, error)
	FindUserById(userId string) (*model.User, error)
	FindUsers() ([]*model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	preQueries := []string{
		`create table users (
			id text,
			login text,
			email text,
			name text,
			surname text,
			password text,
			telegram text
		)`,
	}
	for _, query := range preQueries {
		row := db.QueryRow(query)
		if row.Err() != nil {
			log.Error(row.Err().Error())
		}
	}
	return &userRepository{db}
}

func (ur *userRepository) CreateUser(user *model.User) (*model.User, error) {
	row := ur.db.QueryRow("insert into users (id, login, email, name, surname, password, telegram) "+
		"values ($1, $2, $3, $4, $5, $6, $7)", &user.ID, &user.Login, &user.Email, &user.Name, &user.Surname, &user.Password, &user.Telegram)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return user, nil
}

func (ur *userRepository) FindUserById(userId string) (*model.User, error) {
	user := &model.User{}
	err := ur.db.QueryRow("select * from users where id = $1", &userId).
		Scan(&user.ID, &user.Login, &user.Email, &user.Name, &user.Surname, &user.Password, &user.Telegram)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) FindUsers() ([]*model.User, error) {
	return nil, nil
}

package repository

import (
	"database/sql"
	"fmt"
	"graduation-project/challenge-api/model"

	"github.com/prometheus/common/log"
)

type UserRepository interface {
	CreateUser(user *model.User) (*model.User, error)
	FindUserById(userId string) (*model.User, error)
	FindUsers() ([]*model.User, error)
	FindUserByUsername(login string) (*model.User, error)
	FindUserByTelegram(userTelegram string) (*model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	preQueries := []string{
		`create table users (
			id text,
			username text,
			password text,
			email text,
			telegram text
		)`,
	}
	for _, query := range preQueries {
		row := db.QueryRow(query)
		if row.Err() != nil {
			log.Error(row.Err().Error())
		}
		fmt.Println("\n\n\n===========", row.Err(), row.Scan())
	}
	return &userRepository{db}
}

func (ur *userRepository) CreateUser(user *model.User) (*model.User, error) {
	row := ur.db.QueryRow("insert into users (id, username, password, email, telegram) "+
		"values ($1, $2, $3, $4, $5 )", &user.ID, &user.Username, &user.Password, &user.Email, &user.Telegram)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return user, nil
}

func (ur *userRepository) FindUserById(userId string) (*model.User, error) {
	user := &model.User{}
	err := ur.db.QueryRow("select * from users where id = $1", &userId).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Telegram)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) FindUsers() ([]*model.User, error) {
	return nil, nil
}

func (ur *userRepository) FindUserByUsername(username string) (*model.User, error) {
	user := &model.User{}
	err := ur.db.QueryRow("select * from users where username = $1", &username).
		Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Telegram)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) FindUserByTelegram(userTelegram string) (*model.User, error) {
	fmt.Println(&userTelegram, userTelegram)
	user := &model.User{}
	err := ur.db.QueryRow("select * from users where telegram = $1", &userTelegram).
		Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Telegram)
	fmt.Println("\n\n==== repo 1", user, userTelegram)
	return user, err
}

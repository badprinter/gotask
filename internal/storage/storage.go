package storage

import (
	"database/sql"
	"gotask/internal/config"
	"gotask/internal/models"

	_ "github.com/lib/pq"
)

type Storage struct {
	DB *sql.DB
}

func NewStorage(c *config.DB_Config) (*Storage, error) {

	db, err := sql.Open("postgres", c.GetUrlPostgres())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	stor := &Storage{
		DB: db,
	}
	return stor, nil

}

func (stor *Storage) InsertUser(user *models.User) error {
	if err := stor.DB.QueryRow("INSERT INTO users (user_name, password, email) VALUES ($1, $2, $3) RETURNING user_id",
		user.User_name,
		user.Password,
		user.Email,
	).Scan(&user.ID); err != nil {
		return err
	}
	return nil
}

func (stor *Storage) InserUserByParam(name, password, email string) (*models.User, error) {
	var u models.User
	if err := stor.DB.QueryRow("INSERT INTO users (user_name, password, email) VALUES ($1, $2, $3) RETURNING user_id, user_name, password, email",
		name,
		password,
		email,
	).Scan(&u.ID, &u.User_name, &u.Password, &u.Email); err != nil {
		return nil, err
	}
	return &u, nil
}

func (stor *Storage) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := stor.DB.QueryRow(
		"SELECT  user_id, user_name, password, email FROM users WHERE email = $1", email).Scan(&user.ID, &user.User_name, &user.Password, &user.Email); err != nil {
		return nil, err
	}

	return &user, nil
}

package models

import (
	"database/sql"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int
	Email        string
	PasswordHash string
}

type UserService struct {
	DB *sql.DB
}

func (us *UserService) Create(email, password string) (*User, error) {
	email = strings.ToLower(email)
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("generate password hash: %w", err)
	}

	user := User{
		Email:        email,
		PasswordHash: string(hashBytes),
	}
	row := us.DB.QueryRow(`
		INSERT INTO users (email, password_hash)
		VALUES ($1, $2) RETURNING id`, email, user.PasswordHash)
	err = row.Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("insert data into DB: %w", err)
	}
	return &user, nil
}

func (us *UserService) Authenticate(email, password string) (*User, error) {
	email = strings.ToLower(email)
	user := User{
		Email: email,
	}
	row := us.DB.QueryRow(`
		SELECT id, password_hash
		FROM users WHERE email=$1`, email)
	err := row.Scan(&user.ID, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("password is not valid: %w", err)
	}
	return &user, nil
}

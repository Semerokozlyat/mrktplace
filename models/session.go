package models

import (
	"database/sql"
	"fmt"

	"mrktplace/rand"
)

type Session struct {
	ID        int
	UserID    int
	Token     string // Token is only set when new session is created - but not when it is retrieved from DB.
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	token, err := rand.SessionToken()
	if err != nil {
		return nil, fmt.Errorf("generate token: %w", err)
	}
	// TODO: save session to DB, generate token hash
	s := Session{
		UserID: userID,
		Token:  token,
	}
	return &s, nil
}

func (ss *SessionService) User(sessionToken string) (*User, error) {
	return nil, nil
}

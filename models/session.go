package models

import "database/sql"

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
	// TODO: implement
	return nil, nil
}

func (ss *SessionService) User(sessionToken string) (*User, error) {
	return nil, nil
}

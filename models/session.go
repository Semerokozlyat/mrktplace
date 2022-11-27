package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"

	"mrktplace/rand"
)

const (
	MinBytesPerToken = 32 // The minimum number of bytes used for each session token.
)

type Session struct {
	ID        int
	UserID    int
	Token     string // Token is only set when new session is created - but not when it is retrieved from DB.
	TokenHash string
}

type SessionService struct {
	DB            *sql.DB
	BytesPerToken int
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	bytesPerToken := ss.BytesPerToken
	if bytesPerToken < MinBytesPerToken {
		bytesPerToken = MinBytesPerToken
	}
	token, err := rand.String(bytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("generate token: %w", err)
	}
	s := Session{
		UserID:    userID,
		Token:     token,
		TokenHash: ss.hash(token),
	}
	row := ss.DB.QueryRow(`
	UPDATE sessions
	SET token_hash = $2
	WHERE user_id = $1 
	RETURNING id;`, s.UserID, s.TokenHash)
	err = row.Scan(&s.ID)
	if errors.Is(err, sql.ErrNoRows) {
		row = ss.DB.QueryRow(`
		INSERT INTO sessions (user_id, token_hash)
		VALUES($1, $2) 
		RETURNING id;`, s.UserID, s.TokenHash)
		err = row.Scan(&s.ID)
	}
	if err != nil {
		return nil, fmt.Errorf("insert session data into db: %w", err)
	}
	return &s, nil
}

func (ss *SessionService) User(sessionToken string) (*User, error) {
	tokenHash := ss.hash(sessionToken)
	var user User
	row := ss.DB.QueryRow(`
	SELECT user_id 
	FROM sessions
	WHERE token_hash = $1;`, tokenHash)
	err := row.Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("query user id by session token hash: %w", err)
	}
	row = ss.DB.QueryRow(`
	SELECT email, password_hash
	FROM users
	WHERE id = $1`, user.ID)
	err = row.Scan(&user.Email, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("query user by id: %w", err)
	}
	return &user, nil
}

func (ss *SessionService) hash(sessionToken string) string {
	tokenHash := sha256.Sum256([]byte(sessionToken))
	return base64.URLEncoding.EncodeToString(tokenHash[:])
}

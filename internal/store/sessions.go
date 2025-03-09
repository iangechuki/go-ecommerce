package store

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"time"
)

type Session struct {
	ID                int64     `json:"id"`
	UserID            int64     `json:"user_id"`
	TokenHash         string    `json:"token_hash"`
	CreatedAt         time.Time `json:"created_at"`
	ExpiresAt         time.Time `json:"expires_at"`
	LastAccessed      time.Time `json:"last_accessed"`
	IPAddress         string    `json:"ip_address"`
	UserAgent         string    `json:"user_agent"`
	ClientFingerprint string    `json:"client_fingerprint"`
}

type SessionStore struct {
	db *sql.DB
}

func (s *SessionStore) Create(ctx context.Context, session *Session) error {
	query := `
	INSERT INTO sessions
	(user_id, token_hash, created_at, expires_at, last_accessed, ip_address, user_agent, client_fingerprint)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
	RETURNING id,created_at,last_accessed
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	err := s.db.QueryRowContext(ctx,
		query,
		session.UserID,
		session.TokenHash,
		session.CreatedAt,
		session.ExpiresAt,
		session.LastAccessed,
		session.IPAddress,
		session.UserAgent,
		session.ClientFingerprint,
	).
		Scan(
			&session.ID,
			&session.CreatedAt,
			&session.LastAccessed)

	if err != nil {
		return err
	}
	return nil
}
func (s *SessionStore) GetByUserFingerprint(ctx context.Context, userID int64, fingerprint string) (*Session, error) {
	query := `
	SELECT id, user_id, token_hash, created_at, expires_at, last_accessed, ip_address, user_agent
	FROM sessions
	WHERE user_id = $1 AND client_fingerprint = $2
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	var session Session
	err := s.db.QueryRowContext(ctx, query, userID, fingerprint).
		Scan(&session.ID,
			&session.UserID,
			&session.TokenHash,
			&session.CreatedAt,
			&session.ExpiresAt,
			&session.LastAccessed,
			&session.IPAddress,
			&session.UserAgent)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &session, nil
}
func (s *SessionStore) GetByToken(ctx context.Context, token string) (*Session, error) {
	tokenHash := sha256.Sum256([]byte(token))
	hashString := hex.EncodeToString(tokenHash[:])
	query := `
	SELECT id, user_id, token_hash, created_at, expires_at, last_accessed, ip_address, user_agent, client_fingerprint
	FROM sessions
	WHERE token_hash = $1
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	var session Session
	err := s.db.QueryRowContext(ctx, query, hashString).
		Scan(&session.ID,
			&session.UserID,
			&session.TokenHash,
			&session.CreatedAt,
			&session.ExpiresAt,
			&session.LastAccessed,
			&session.IPAddress,
			&session.UserAgent,
			&session.ClientFingerprint,
		)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &session, nil
}
func (s *SessionStore) UpdateLastAccessed(ctx context.Context, sessionID int64) error {
	query := `
	UPDATE sessions
	SET last_accessed = NOW()
	WHERE id = $1
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	_, err := s.db.ExecContext(ctx, query, sessionID)
	if err != nil {
		return err
	}
	return nil
}
func (s *SessionStore) Delete(ctx context.Context, sessionID int64) error {
	query := `
	DELETE FROM sessions
	WHERE id = $1
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	_, err := s.db.ExecContext(ctx, query, sessionID)
	if err != nil {
		return err
	}
	return nil
}
func (s *SessionStore) DeleteExpired(ctx context.Context) (result sql.Result, err error) {
	query := `
	DELETE FROM SESSIONS
	WHERE expires_at < NOW()
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	result, err = s.db.ExecContext(ctx, query)

	return result, err
}
func (s *SessionStore) DeleteByUserID(ctx context.Context, userID int64) error {
	query := `
	DELETE FROM sessions
	WHERE user_id = $1
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	_, err := s.db.ExecContext(ctx, query, userID)
	if err != nil {
		return err
	}
	return nil
}

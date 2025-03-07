package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	QueryTimeoutDuration = time.Second * 5
	ErrRecordNotFound    = errors.New("record not found")
)

type Storage struct {
	Users interface {
		Create(context.Context, *User) error
		CreateAndInvite(ctx context.Context, user *User, token string, expiresIn time.Duration) error
		GetByEmail(context.Context, string) (*User, error)
		GetByID(context.Context, int64) (*User, error)
		Update(context.Context, *User) error
		Delete(context.Context, int64) error
		Verify(ctx context.Context, token string) error
		CreatePasswordResetToken(ctx context.Context, user *User, token string, expiresIn time.Duration) error
		ResetPasswordUsingToken(ctx context.Context, token string, newPassword string) error
	}
	Sessions interface {
		Create(context.Context, *Session) error
		GetByToken(ctx context.Context, token string) (*Session, error)
		// GetByTokenID(context.Context, string) (*Session, error)
		Delete(context.Context, int64) error
		UpdateLastAccessed(context.Context, int64) error
		DeleteExpired(context.Context) error
		DeleteByUserID(ctx context.Context, userID int64) error
	}
}

func NewPostgresStorage(db *sql.DB) Storage {
	return Storage{
		Users:    &UserStore{db: db},
		Sessions: &SessionStore{db: db},
	}
}
func withTx(db *sql.DB, ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	if err := fn(tx); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}
	return tx.Commit()
}

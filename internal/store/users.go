package store

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicateEmail   = errors.New("user with email already exists")
	ErrPasswordRequired = errors.New("password is required")
)

type User struct {
	ID           int64     `json:"id"`
	Email        string    `json:"email"`
	Username     string    `json:"username"`
	Password     password  `json:"-"`
	IsVerified   bool      `json:"is_verified"`
	TwoFAEnabled bool      `json:"two_fa_enabled"`
	TwoFASecret  string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
type password struct {
	text string
	hash []byte
}

func (p *password) Set(text string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	p.text = text
	p.hash = hash
	return nil
}
func (p *password) Compare(text string) error {
	return bcrypt.CompareHashAndPassword(p.hash, []byte(text))
}

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, u *User) error {
	query := `
		INSERT INTO users
		(email,username,password_hash,is_verified,two_fa_enabled,created_at,updated_at)
		VALUES ($1,$2,$3,$4,$5,NOW(),NOW())
		RETURNING id,created_at,updated_at
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		u.Email,
		u.Username,
		u.Password.hash,
		u.IsVerified,
		u.TwoFAEnabled,
	).
		Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		case err.Error() == `pq: null value in column "password_hash" violates not-null constraint`:
			return ErrPasswordRequired
		default:
			return err
		}
	}

	return nil
}
func (s *UserStore) CreateAndInvite(ctx context.Context, user *User, token string, expiresIn time.Duration) error {
	return withTx(s.db, ctx, func(tx *sql.Tx) error {
		if err := s.Create(ctx, user); err != nil {
			return err
		}
		if err := s.createAccountVerificationToken(ctx, tx, user.ID, token, expiresIn); err != nil {
			return err
		}
		return nil
	})
}
func (s *UserStore) GetByEmail(ctx context.Context, email string) (*User, error) {
	query := `
	SELECT id,email,username,password_hash,is_verified,two_fa_enabled,two_fa_secret,created_at,updated_at
	FROM users
	WHERE email = $1
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var user User
	err := s.db.QueryRowContext(ctx, query, email).
		Scan(&user.ID,
			&user.Email,
			&user.Username,
			&user.Password.hash,
			&user.IsVerified,
			&user.TwoFAEnabled,
			&user.TwoFASecret,
			&user.CreatedAt,
			&user.UpdatedAt)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &user, nil
}
func (s *UserStore) GetByID(ctx context.Context, id int64) (*User, error) {
	query := `
	SELECT id,email,username,password_hash,is_verified,two_fa_enabled,two_fa_secret,created_at,updated_at
	FROM users
	WHERE id = $1
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	var user User
	err := s.db.QueryRowContext(ctx, query, id).
		Scan(&user.ID,
			&user.Email,
			&user.Username,
			&user.Password.hash,
			&user.IsVerified,
			&user.TwoFAEnabled,
			&user.TwoFASecret,
			&user.CreatedAt,
			&user.UpdatedAt)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &user, nil
}
func (s *UserStore) Update(ctx context.Context, u *User) error {
	query := `
	UPDATE users
	SET email = $1, is_verified = $2, two_fa_enabled = $3, two_fa_secret = $4, updated_at = NOW()
	WHERE id = $5
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	_, err := s.db.ExecContext(ctx, query, u.Email, u.IsVerified, u.TwoFAEnabled, u.TwoFASecret, u.ID)
	if err != nil {
		return err
	}
	return nil
}
func (s *UserStore) Delete(ctx context.Context, id int64) error {
	query := `
	DELETE FROM users
	WHERE id = $1
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
func (s *UserStore) Verify(ctx context.Context, token string) error {
	return withTx(s.db, ctx, func(tx *sql.Tx) error {
		// 1.Get the user that the token belongs to
		user, err := s.getUserFromVerificationToken(ctx, tx, token)
		if err != nil {
			return err
		}
		// 2.update the user to verified
		user.IsVerified = true
		if err := s.Update(ctx, user); err != nil {
			return err
		}
		// 3.delete the account verification token
		if err := s.deleteAccountVerificationToken(ctx, tx, user.ID); err != nil {
			return err
		}
		return nil
	})
}
func (s *UserStore) deleteAccountVerificationToken(ctx context.Context, tx *sql.Tx, userID int64) error {
	query := `
	DELETE FROM account_verification_tokens
	WHERE user_id = $1
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	_, err := tx.ExecContext(ctx, query, userID)
	if err != nil {
		return err
	}
	return nil
}
func (s *UserStore) createAccountVerificationToken(ctx context.Context, tx *sql.Tx, userID int64, token string, expiresIn time.Duration) error {
	query := `
	INSERT INTO 
	account_verification_tokens (user_id, token_hash, expires_at)
	VALUES ($1, $2, $3)
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	_, err := tx.ExecContext(ctx, query, userID, token, time.Now().Add(expiresIn))
	if err != nil {
		return err
	}
	return nil
}

func (s *UserStore) getUserFromVerificationToken(ctx context.Context, tx *sql.Tx, token string) (*User, error) {
	query := `
	SELECT u.id,u.username,u.email,u.created_at,u.is_verified
	FROM users u
	JOIN account_verification_tokens at ON u.id = at.user_id
	WHERE at.token_hash = $1 AND at.expires_at > $2
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	hash := sha256.Sum256([]byte(token))
	hashToken := hex.EncodeToString(hash[:])
	var user User
	err := tx.QueryRowContext(ctx, query, hashToken, time.Now()).
		Scan(&user.ID,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
			&user.IsVerified,
		)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &user, nil
}

// Password
func (s *UserStore) CreatePasswordResetToken(ctx context.Context, user *User, token string, expiresIn time.Duration) error {
	query := `
	INSERT INTO password_reset_tokens (user_id, token_hash, expires_at)
	VALUES ($1, $2, $3)
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	_, err := s.db.ExecContext(ctx, query, user.ID, token, time.Now().Add(expiresIn))
	if err != nil {
		return err
	}
	return nil
}
func (u *UserStore) getUserFromPasswordResetToken(ctx context.Context, token string) (*User, error) {
	query := `
	SELECT u.id,u.username,u.email,u.created_at,u.is_verified
	FROM users u
	JOIN password_reset_tokens prt ON u.id = prt.user_id
	WHERE prt.token_hash = $1 AND prt.expires_at > $2
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	hash := sha256.Sum256([]byte(token))
	hashToken := hex.EncodeToString(hash[:])
	var user User
	err := u.db.QueryRowContext(ctx, query, hashToken, time.Now()).
		Scan(&user.ID,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
			&user.IsVerified,
		)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &user, nil
}
func (u *UserStore) deleteAllSessions(ctx context.Context, userID int64) error {
	query := `
	DELETE FROM sessions
	WHERE user_id = $1
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	_, err := u.db.ExecContext(ctx, query, userID)
	if err != nil {
		return err
	}
	return nil
}
func (u *UserStore) deletePasswordResetToken(ctx context.Context, user *User) error {
	query := `
	DELETE FROM password_reset_tokens
	WHERE user_id = $1
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	_, err := u.db.ExecContext(ctx, query, user.ID)
	if err != nil {
		return err
	}
	return nil
}
func (u *UserStore) UpdatePassword(ctx context.Context, user *User, password string) error {
	query := `
	UPDATE users
	SET password_hash = $2
	WHERE id = $1
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := u.db.ExecContext(ctx, query, user.ID, password)
	if err != nil {
		return err
	}
	return nil
}
func (u *UserStore) ResetPasswordUsingToken(ctx context.Context, token string, newPassword string) error {
	return withTx(u.db, ctx, func(tx *sql.Tx) error {
		user, err := u.getUserFromPasswordResetToken(ctx, token)
		if err != nil {
			return err
		}
		if err := user.Password.Set(newPassword); err != nil {
			return err
		}
		if err := u.UpdatePassword(ctx, user, string(user.Password.hash)); err != nil {
			return err
		}
		if err := u.deletePasswordResetToken(ctx, user); err != nil {
			return err
		}
		if err := u.deleteAllSessions(ctx, user.ID); err != nil {
			return err
		}
		return nil
	})
}

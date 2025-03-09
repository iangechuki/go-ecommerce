package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserID int64  `json:"user_id"`
	Type   string `json:"type"`
	jwt.RegisteredClaims
}

//	func (app *application) GenerateVerificationToken(userID int64) (string, error) {
//		claims := CustomClaims{
//			UserID: userID,
//			Type:   "verify",
//			RegisteredClaims: jwt.RegisteredClaims{
//				Issuer:    app.config.auth.iss,
//				Audience:  jwt.ClaimStrings{app.config.auth.aud},
//				ExpiresAt: jwt.NewNumericDate(time.Now().Add(app.config.auth.accessToken.exp)),
//			},
//		}
//		tokenString, err := app.accessAuthenticator.GenerateToken(claims)
//		if err != nil {
//			return "", err
//		}
//		return tokenString, nil
//	}
//
//	func (app *application) GenerateResetPasswordToken(userID int64) (string, error) {
//		claims := CustomClaims{
//			UserID: userID,
//			Type:   "reset",
//			RegisteredClaims: jwt.RegisteredClaims{
//				Issuer:    app.config.auth.iss,
//				Audience:  jwt.ClaimStrings{app.config.auth.aud},
//				ExpiresAt: jwt.NewNumericDate(time.Now().Add(app.config.auth.accessToken.exp)),
//			},
//		}
//		tokenString, err := app.accessAuthenticator.GenerateToken(claims)
//		if err != nil {
//			return "", err
//		}
//		return tokenString, nil
//	}
func (app *application) GenerateAccessToken(userID int64) (string, error) {
	claims := CustomClaims{
		UserID: userID,
		Type:   "access",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    app.config.auth.iss,
			Audience:  jwt.ClaimStrings{app.config.auth.aud},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(app.config.auth.accessToken.exp)),
		},
	}
	tokenString, err := app.accessAuthenticator.GenerateToken(claims)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func (app *application) GenerateRefreshToken(userID int64) (string, error) {
	claims := CustomClaims{
		UserID: userID,
		Type:   "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    app.config.auth.iss,
			Audience:  jwt.ClaimStrings{app.config.auth.aud},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(app.config.auth.refreshToken.exp)),
		},
	}
	tokenString, err := app.refreshAuthenticator.GenerateToken(claims)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func (app *application) Generate2FAToken(userID int64) (string, error) {
	claims := CustomClaims{
		UserID: userID,
		Type:   "2fa",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    app.config.auth.iss,
			Audience:  jwt.ClaimStrings{app.config.auth.aud},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
		},
	}
	token, err := app.accessAuthenticator.GenerateToken(claims)
	if err != nil {
		return "", err
	}
	return token, nil

}
func (app *application) CreateClientFingerprint(r *http.Request) string {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	if ip == "" {
		ip = r.RemoteAddr
	}
	if strings.Contains(ip, "::1") {
		ip = "127.0.0.1"
	}
	// combine IP and User-Agent
	fingerprint := fmt.Sprintf("%s-%s", ip, r.UserAgent())
	hash := sha256.Sum256([]byte(fingerprint))
	return hex.EncodeToString(hash[:])
}

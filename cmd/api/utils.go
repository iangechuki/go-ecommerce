package main

import (
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

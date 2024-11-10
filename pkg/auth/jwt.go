package auth

import (
    "time"
    "github.com/golang-jwt/jwt/v5"
)

type JWTCustomClaims struct {
    UserID uint `json:"userId"`
    jwt.RegisteredClaims
}

func GenerateToken(userID uint, secret string) (string, error) {
    claims := &JWTCustomClaims{
        userID,
        jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secret))
}

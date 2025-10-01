package utils

import (
    "errors"
    "github.com/golang-jwt/jwt"
    "strconv"
    "time"
)

var jwtKey = []byte("secret_key") // replace with secure env variable in production

// GenerateJWT creates a token for a user ID
func GenerateJWT(userID uint) (string, error) {
    claims := &jwt.StandardClaims{
        ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
        Subject:   strconv.Itoa(int(userID)),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

// ValidateJWT parses and validates a token, returns userID
func ValidateJWT(tokenStr string) (uint, error) {
    token, err := jwt.ParseWithClaims(tokenStr, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if err != nil {
        return 0, err
    }

    if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
        id, err := strconv.Atoi(claims.Subject)
        if err != nil {
            return 0, err
        }
        return uint(id), nil
    }
    return 0, errors.New("invalid token")
}


package jwtx

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string
	jwt.RegisteredClaims
}

type TokenInfo struct {
	ExpiresAt   time.Time              `json:"expiresAt"`
	TokenType   *jwt.SigningMethodHMAC `json:"tokenType"`
	AccessToken string                 `json:"accessToken"`
}

func GenerateToken(subject string) (*TokenInfo, error) {
	now := time.Now()
	expiresAt := now.Add(time.Duration(7200) * time.Second)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		NotBefore: jwt.NewNumericDate(now),
		Subject:   subject,
	})

	tokenString, err := token.SignedString([]byte("qinglv"))
	if err != nil {
		fmt.Println("GenerateToken", subject, err)
		return nil, err
	}

	tokenInfo := &TokenInfo{
		ExpiresAt:   expiresAt,
		TokenType:   jwt.SigningMethodHS256,
		AccessToken: tokenString,
	}

	return tokenInfo, nil
}

func ParseToken(tokenString string) (*Claims, error) {
	if tokenString == "" {
		return nil, errors.New("token is empty")
	}
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("qinglv"), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, nil
}

func GetToken(c *gin.Context) string {
	var token string
	auth := c.GetHeader("Authorization")
	prefix := "Bearer "
	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	}
	return token

}

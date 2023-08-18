package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gitnub.com/artemKapitonov/libraryAPI/internal/models"
)

const (
	salt       = "qwsdfghjkoplmnbvcxz14gr7852qazsqw"
	signingKey = "1qr84kjw88qefwrgrg5vgkmfdg4gadsg"
	tokenTTL   = 12 * time.Hour
)

type AuthorizationPostgres interface {
	NewUser(user models.User) (int, error)
}

func (s *Service) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	return s.Repository.NewUser(user)
}

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *Service) GenerateToken(username, password string) (string, error) {
	user, err := s.Repository.User(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *Service) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type")
	}

	return claims.UserID, nil

}

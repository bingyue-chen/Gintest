package services

import (
	"Gintest/src/entities"
	"Gintest/src/repositories"
	"Gintest/src/utilities"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

type AuthenticationServiceContract interface {
	Authenticate(*entities.UserCredential) (*entities.User, bool)
	GenerateToken(*entities.User) (string, error)
}

type AuthenticationService struct {
	UserRepository repositories.UserRepositoryContract
}

var jwtSecret = []byte(utilities.Getenv("JWT_SECRET"))

func (service *AuthenticationService) Authenticate(credential *entities.UserCredential) (*entities.User, bool) {

	var user *entities.User
	var err error

	user, err = service.UserRepository.FindByEmail(credential.Email)

	if err == nil {

		hash_utility := utilities.Hash{}

		return user, hash_utility.Check(user.Password, credential.Password)
	}

	return nil, false
}

func (service *AuthenticationService) GenerateToken(user *entities.User) (string, error) {

	now := time.Now()
	userId := strconv.FormatInt(user.ID, 10)
	id := userId + strconv.FormatInt(now.Unix(), 10)

	claims := entities.Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			Audience:  userId,
			ExpiresAt: now.Add(7 * 24 * time.Hour).Unix(),
			Id:        id,
			IssuedAt:  now.Unix(),
			Issuer:    utilities.Getenv("APP_NAME"),
			NotBefore: now.Add(1 * time.Second).Unix(),
			Subject:   userId,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return tokenClaims.SignedString(jwtSecret)
}

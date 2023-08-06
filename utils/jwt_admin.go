package utils

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// JWTSecretKey is the secret key used to sign JWT tokens. This should be kept secure and should not be hard-coded in production.

var ErrInvalidCredentialsAdmin = errors.New("invalid credentials")

type SignedAdminDetails struct {
    Aid        uint
    Email   string
    jwt.StandardClaims
}

// GenerateToken generates a new JWT token with the given user ID as the subject.
func GenerateTokenAdmin(adminID uint, email string) (signedToken string, err error) {
	
	claims := &SignedAdminDetails{
        Aid:    adminID,
        Email:  email,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
        },
    }
	
 
	secretKey, err := GetJWTSecretKey()
	if err != nil {
		log.Panic(err)
		return
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secretKey)
    
	if err != nil {
		log.Panic(err)
		return
	}
	
    return token, err
}

// ValidateToken validates the provided JWT token and returns the token object if it's valid.
func ValidateTokenAdmin(signedToken string) (claims *SignedAdminDetails, err error) {
    token, err := jwt.ParseWithClaims(
        signedToken,
        &SignedDetails{},
        func(token *jwt.Token) (interface{}, error) {
            return GetJWTSecretKey()
        },
    )

    if err != nil {
        return nil, err
    }
 
    claims, ok := token.Claims.(*SignedAdminDetails)
    if !ok {
		
        return nil, fmt.Errorf("the token is invalid")
    }
 
    if claims.ExpiresAt < time.Now().Local().Unix() {
        return nil, fmt.Errorf("token is expired")
    }
    return claims, nil
}
// HashPassword hashes the provided password using a secure hashing algorithm (bcrypt).
func HashPasswordAdmin(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPasswordHash compares the provided password with the hashed password to check if they match.
func CheckPasswordHashAdmin(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
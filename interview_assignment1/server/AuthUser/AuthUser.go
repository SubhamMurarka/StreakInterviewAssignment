package AuthUser

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtManager struct {
	secretKey     string
	tokenDuration time.Duration
	invalidToken  map[string]bool
	mu            sync.RWMutex
}

type TokenParams struct {
	jwt.StandardClaims
	UserName string
}

type JwtManagerInterface interface {
	GenerateToken(username string) (string, error)
	VerifyToken(accesstoken string) (*TokenParams, error)
	Logout(Accesstoken string) error
}

func NewJwtManager(secretKey string, duration time.Duration) JwtManagerInterface {
	return &JwtManager{
		secretKey:     secretKey,
		tokenDuration: duration,
		invalidToken:  make(map[string]bool),
	}
}

func (j *JwtManager) GenerateToken(username string) (string, error) {
	claims := &TokenParams{
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(j.tokenDuration).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}

	j.mu.Lock()
	defer j.mu.Unlock()

	j.invalidToken[token] = false

	return token, nil
}

func (j *JwtManager) VerifyToken(accesstoken string) (*TokenParams, error) {
	accestoken, err := jwt.ParseWithClaims(
		accesstoken,
		&TokenParams{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.secretKey), nil
		},
	)

	if err != nil {
		log.Printf("Token not able to veriy %v", err)
		return nil, fmt.Errorf("Token not able to veriy %v", err)
	}

	claims, ok := accestoken.Claims.(*TokenParams)
	if !ok {
		log.Printf("Token not able to Fetch claims %v", err)
		return nil, fmt.Errorf("Token not able to Fetch claims %v", err)
	}

	if j.invalidToken[accesstoken] == false {
		return nil, fmt.Errorf("User Logged Out")
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, fmt.Errorf("Token Expired")
	}

	return claims, nil
}

func (j *JwtManager) Logout(Accesstoken string) error {
	j.mu.Lock()
	defer j.mu.Unlock()

	if j.invalidToken[Accesstoken] == true {
		return fmt.Errorf("Not Allowed")
	}

	j.invalidToken[Accesstoken] = true

	return nil
}

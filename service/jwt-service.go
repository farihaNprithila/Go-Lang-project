package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

/*JWT service is a contract what JWT can do*/
type JWTService interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

/*NewJWTService method is to create a new instance of JWTService*/
func NewJWTService() *jwtService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "root",
	}
}
func getSecretKey() string {
	secretKey := os.Getenv("JWT_KEY")
	if secretKey != "" {
		secretKey = "root"
	}
	return secretKey
}

/*For generating the secret key*/
func (j *jwtService) GenerateToken(userID string) string {
	claims := &jwtCustomClaim{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 0, 1).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	/*Token Generation*/
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

/*Validation of the token*/
func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpeted siging method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})

}

package pkg

/*
import (
	"errors"
	"time"
	"user-service/internal/domain"

	"github.com/dgrijalva/jwt-go"
)

// JWTClaims represents the claims stored in the JWT token.
type JWTClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

var (
	jwtSecretKey      = "todo-microservice-nullbyte"
	jwtExpirationTime = time.Hour * 24
)

// generateJWT creates a new JWT token for the user.
func generateJWT(user *domain.User) (string, error) {
	claims := &JWTClaims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(jwtExpirationTime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}

// validateJWT parses and validates a JWT token.
func validateJWT(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecretKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
*/

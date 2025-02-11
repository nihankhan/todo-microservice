package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("todo-microservice")

// AuthMiddleware checks if the incoming request has a valid JWT token
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var tokenString string

		// First, try to extract the token from the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader != "" {
			parts := strings.Split(authHeader, " ")
			if len(parts) == 2 && parts[0] == "Bearer" {
				tokenString = parts[1]
			}
		}

		// If the token is not in the Authorization header, check the Cookie
		if tokenString == "" {
			cookie, err := r.Cookie("access_token")
			if err == nil {
				tokenString = cookie.Value
			}
		}

		// If no token is found, return unauthorized error
		if tokenString == "" {
			http.Error(w, "Missing Authorization or Cookie token", http.StatusUnauthorized)
			return
		}

		// Parse and validate the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Ensure the signing method is correct
			if token.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Pass the request to the next handler, with the user info
		ctx := r.Context()
		ctx = context.WithValue(ctx, "user_id", token.Claims.(jwt.MapClaims)["sub"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// func AuthMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		token := r.Header.Get("Authorization")
// 		if token == "" || !strings.HasPrefix(token, "Bearer") {
// 			http.Error(w, "Forbidden", http.StatusForbidden)
// 			return
// 		}

// 		// token = token[len("Bearer "):]
// 		// if token !=

// 		next.ServeHTTP(w, r)
// 	})
// }

// Simple logging middleware
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

/*

package middleware

import (
    "fmt"
    "net/http"
    "strings"
    "github.com/dgrijalva/jwt-go"
    "context"
)

// Define the secret key for JWT (this can be moved to an environment variable)
var secretKey = []byte("your-secret-key")

// AuthMiddleware checks if the incoming request has a valid JWT token
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        var tokenString string

        // First, try to extract the token from the Authorization header
        authHeader := r.Header.Get("Authorization")
        if authHeader != "" {
            parts := strings.Split(authHeader, " ")
            if len(parts) == 2 && parts[0] == "Bearer" {
                tokenString = parts[1]
            }
        }

        // If the token is not in the Authorization header, check the Cookie
        if tokenString == "" {
            cookie, err := r.Cookie("access_token")
            if err == nil {
                tokenString = cookie.Value
            }
        }

        // If no token is found, return unauthorized error
        if tokenString == "" {
            http.Error(w, "Missing Authorization or Cookie token", http.StatusUnauthorized)
            return
        }

        // Parse and validate the JWT token
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHS256); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return secretKey, nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // Pass the request to the next handler, with the user info
        ctx := r.Context()
        ctx = context.WithValue(ctx, "user", token.Claims.(jwt.MapClaims)["sub"])
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}


*/

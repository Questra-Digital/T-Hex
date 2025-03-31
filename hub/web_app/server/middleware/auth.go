package middleware

import (
	"net/http"
	"server/models"
	"github.com/golang-jwt/jwt/v4"
	"server/utils"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Extract the token from the Cookie
		cookie, err := r.Cookie("token")
		if err != nil {
			utils.RespondError(w, "Missing Token", http.StatusUnauthorized)
			return
		}

		// Parse and validate the token
		claims := &models.Claims{}
		token, err := jwt.ParseWithClaims(cookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
			return utils.Jwtsecret, nil
		})

		if err != nil || !token.Valid {
			utils.RespondError(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		// Validate username
		if claims.Username == "" {
			utils.RespondError(w, "Missing Username", http.StatusUnauthorized)
			return
		}

		// Check if username exists in the ValidUsers map
		if _, exists := models.ValidUsers[claims.Username]; !exists {
			utils.RespondError(w, "User does not exist", http.StatusUnauthorized)
			return
		}

		r.Header.Set("Username", claims.Username)

		// Call the next handler
		next(w, r)
	}
}

package middleware

import(
	"net/http"
	"server/models"
	"github.com/golang-jwt/jwt/v4"
	"server/utils"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){

		//Extract token from Authorization Header Field
		tokenString := r.Header.Get("Authorization")

		if tokenString == ""{
			http.Error(w,"Missing Token",http.StatusUnauthorized)
			return
		}

		//Parse and validate the token
		claims := &models.Claims{}
		token,err := jwt.ParseWithClaims(tokenString,claims,func (token *jwt.Token) (interface{},error){
			return utils.Jwtsecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}


		//Set request header for subsequent handlers
		r.Header.Set("Username",claims.Username)

		//Call the next handler
		next(w,r)

	}
}

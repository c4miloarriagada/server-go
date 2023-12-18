package middleware

import (
	"fmt"
	"net/http"
	"os"
	"serverpackage/internal/auth"

	"github.com/joho/godotenv"
)

var (
	redirectURL string
	port        string
)

func init() {
	err := godotenv.Load("")
	if err != nil {
		fmt.Println("Error loading .env file")

	}

	redirectURL = os.Getenv("REDIRECT_URL")
	port = os.Getenv("PORT")
}

func JwtHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("jwt_token")

		if authorizationHeader == "" {
			http.Error(w, "Missing Token", http.StatusUnauthorized)
			redirectWithError(w, r, "Missing Token")
			return
		}

		err := auth.VerifyToken(authorizationHeader)
		if err != nil {
			http.Error(w, "Token Verification Failed", http.StatusUnauthorized)
			redirectWithError(w, r, "Token Verification Failed")
			return
		}

		fmt.Println("Passed token verification")
		next(w, r)
	}
}

func redirectWithError(w http.ResponseWriter, r *http.Request, errorMsg string) {
	redirectURLWithError := fmt.Sprintf("%s/%s?error=%s", redirectURL, port, errorMsg)
	http.Redirect(w, r, redirectURLWithError, http.StatusFound)
}

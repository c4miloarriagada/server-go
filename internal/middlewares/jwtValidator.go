package middlewares

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

func JwtHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Error loading .env file")

		}

		redirectURL = os.Getenv("REDIRECT_URL")
		port = os.Getenv("PORT")

		authorizationHeader := r.Header.Get("jwt_token")

		if authorizationHeader == "" {
			http.Error(w, "Missing Token", http.StatusUnauthorized)
			redirectWithError(w, r, "Missing Token")
			return
		}

		authErr := auth.VerifyToken(authorizationHeader)
		if authErr != nil {
			http.Error(w, "Token Verification Failed", http.StatusUnauthorized)
			redirectWithError(w, r, "Token Verification Failed")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func redirectWithError(w http.ResponseWriter, r *http.Request, errorMsg string) {
	redirectURLWithError := fmt.Sprintf("%s/%s?error=%s", redirectURL, port, errorMsg)
	http.Redirect(w, r, redirectURLWithError, http.StatusFound)
}

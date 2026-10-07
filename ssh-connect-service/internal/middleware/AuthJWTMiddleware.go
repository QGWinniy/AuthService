package middleware

import (
	"AuthService/ssh-connect-service/internal/domain"
	"AuthService/ssh-connect-service/pkg/jwt"
	"context"
	"log"
	"net/http"
	"time"
)

type contextKey string

func AuthJWTMiddleware(next http.HandlerFunc, managerJWT *jwt.Manager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("========== AUTH MIDDLEWARE ==========")
		log.Printf("method=%s path=%s remote=%s", r.Method, r.URL.Path, r.RemoteAddr)

		// Проверяем JWT manager
		if managerJWT == nil {
			log.Println("ERROR: managerJWT == nil")
			http.Error(w, "jwt manager is nil", http.StatusInternalServerError)
			return
		}

		log.Println("JWT manager: OK")

		// Получаем cookie
		cookie, err := r.Cookie(domain.CookieSessionJWT)
		if err != nil {
			log.Printf("COOKIE ERROR: %v", err)
			log.Printf("expected cookie name: %s", domain.CookieSessionJWT)

			http.Error(w, "token expired", http.StatusUnauthorized)
			return
		}

		log.Printf("COOKIE FOUND: name=%s", cookie.Name)
		log.Printf("COOKIE VALUE: %s", cookie.Value)

		if cookie.Value == "" {
			log.Println("ERROR: cookie value is empty")

			http.Error(w, "empty token", http.StatusUnauthorized)
			return
		}

		// Парсим JWT
		log.Println("Parsing JWT...")

		claims, err := managerJWT.Parse(cookie.Value)

		if err != nil {
			log.Printf("JWT PARSE ERROR: %v", err)

			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		log.Println("JWT parsed successfully")

		// Проверяем claims
		if claims == nil {
			log.Println("ERROR: claims == nil")

			http.Error(w, "invalid claims", http.StatusUnauthorized)
			return
		}

		log.Printf("JWT CLAIMS: %+v", claims)

		// Проверяем expiration
		if claims.ExpiresAt == nil {
			log.Println("ERROR: claims.ExpiresAt == nil")

			http.Error(w, "token has no expiration", http.StatusUnauthorized)
			return
		}

		now := time.Now()
		exp := claims.ExpiresAt.Time

		log.Printf("JWT EXP: %v", exp)
		log.Printf("CURRENT TIME: %v", now)
		log.Printf("TOKEN EXPIRED: %v", exp.Before(now))

		if exp.Before(now) {
			log.Println("ERROR: JWT token expired")

			http.SetCookie(w, &http.Cookie{
				Name:     domain.CookieSessionJWT,
				Value:    "",
				Path:     "/",
				MaxAge:   -1,
				HttpOnly: true,
				Secure:   true,
				SameSite: http.SameSiteLaxMode,
			})

			http.Error(w, "token expired", http.StatusUnauthorized)
			return
		}

		// User ID
		log.Printf("JWT USER ID: %v", claims.UserID)

		// Передаём userID в context
		ctx := context.WithValue(
			r.Context(),
			domain.UserIDKeyContext,
			claims.UserID,
		)

		log.Println("JWT authentication successful")
		log.Println("====================================")

		next(w, r.WithContext(ctx))
	}
}
package middleware

import (
	"log"
	"net/http"
)

type IMiddleware interface {
	Middleware(next http.Handler) http.Handler
}

// Define our struct
type authenticationMiddleware struct {
	tokenUsers map[string]string
}

// New Initialize it somewhere
func New(tokenUsers map[string]string) IMiddleware {
	authUsers := make(map[string]string, len(tokenUsers))
	for k, v := range tokenUsers {
		authUsers[v] = k
	}
	return &authenticationMiddleware{tokenUsers: authUsers}
}

// Middleware function, which will be called for each request
func (amw *authenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Session-Token")

		if user, found := amw.tokenUsers[token]; found {
			// We found the token in our map
			log.Printf("Authenticated user %s\n", user)
			// Pass down the request to the next middleware (or final handler)
			next.ServeHTTP(w, r)
		} else {
			// Write an error and stop the handler chain
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

package bcmiddleware

import (
	"fmt"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/jwt"
)

var nonAuthHTML = `
<html>
<body>
<h1>Not Authenticated</h1>
</body>
</html>
`

func SetNonAuthHTML(html string) {
	nonAuthHTML = html
}

// Authenticator is a Chi JWT Auth Middleware
// Use:
//  var jwtAuth = jwtauth.New("HS256", []byte("secret"), nil)
//  r.Group(func(r chi.Router) {
// 		r.Use(jwtauth.Verifier(jwtAuth))
// 		r.Use(bcmiddleware.Authenticator)
//      ... rest of the routes
//  })
func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, err := jwtauth.FromContext(r.Context())

		if err != nil {
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, nonAuthHTML)
			return
		}

		if token == nil || jwt.Validate(token) != nil {
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, nonAuthHTML)
			return
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}

package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"schoolsystem/learninghub-api/authorization"
	"strings"

	"github.com/casbin/casbin"
)

func Authorizer(e *casbin.Enforcer) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			method := r.Method
			path := r.URL.Path
			log.Print(path)

			// allow public paths
			if strings.HasPrefix(path, "/p") {
				next.ServeHTTP(w, r)
			} else {
				authHeader := r.Header.Get("Authorization")

				if authHeader == "" {
					http.Error(w, fmt.Sprintf("Authorization header not found"), http.StatusUnauthorized)
					return
				}
				role, err := authorization.GetRoleFromToken(authHeader)
				if err != nil {
					log.Println(err.Error())
					http.Error(w, fmt.Sprintf("Can't get role from token."), http.StatusUnauthorized)
					return
				}

				// allow users that match the role
				if e.Enforce(role, path, method) {
					next.ServeHTTP(w, r)
				} else {
					http.Error(w, fmt.Sprintf("%s forbidded for %s", path, role), http.StatusUnauthorized)
				}
			}

		}

		return http.HandlerFunc(fn)
	}
}

package middleware

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/yanndr/GoWebApp/security"
)

func AuthHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		if security.GetInstance().ReadCookie(r) {

			next.ServeHTTP(w, r)
		} else {
			returnUrl := base64.URLEncoding.EncodeToString([]byte(r.RequestURI))

			if returnUrl == "" {
				http.Redirect(w, r, "/Account/Login", 307)
			} else {
				http.Redirect(w, r, fmt.Sprintf("/Account/Login?returnUrl=%s", returnUrl), 307)
			}
		}
	}

	return http.HandlerFunc(fn)
}

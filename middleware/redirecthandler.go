package middleware

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func RedirectHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		session, err := store.Get(r, "session-name")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		if r.URL.Query()["returnUrl"] != nil {
			returnUrl := r.URL.Query()["returnUrl"][0]
			log.Printf("session returnUrl:%s", returnUrl)
			session.Values["returnUrl"] = returnUrl
			session.Save(r, w)
		}

		next.ServeHTTP(w, r)
	})
}

func PostLoginRedirectHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		session, err := store.Get(r, "session-name")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		result := session.Values["returnUrl"].(string)
		if result != "" {
			log.Printf("session result:%s", result)
		}

		next.ServeHTTP(w, r)
	})
}

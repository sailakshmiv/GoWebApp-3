package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/yanndr/GoWebApp/security"
)

var cookieCodec *securecookie.SecureCookie

var hashKey = securecookie.GenerateRandomKey(64)
var blockKey = securecookie.GenerateRandomKey(64)

func init() {
	cookieCodec = securecookie.New([]byte(hashKey), []byte(blockKey))
}

func LoggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}

	return http.HandlerFunc(fn)
}

func RecoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func AuthHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("AuthHandler")
		if security.GetInstance().ReadCookie(r) {

			next.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, "/Account/Login", 307)
		}
	}

	return http.HandlerFunc(fn)
}

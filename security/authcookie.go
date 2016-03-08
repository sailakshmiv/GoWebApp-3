package security

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/securecookie"
)

type AuthCookie struct {
}

var instance *AuthCookie
var once sync.Once
var hashKey, blockKey []byte
var cookieCodec *securecookie.SecureCookie

const cookieName = "_authCookie"

func GetInstance() *AuthCookie {
	once.Do(func() {
		instance = newAuthCookie()
	})
	return instance
}

func newAuthCookie() *AuthCookie {

	a := &AuthCookie{}
	hashKey = securecookie.GenerateRandomKey(64)
	blockKey = securecookie.GenerateRandomKey(32)
	cookieCodec = securecookie.New(hashKey, blockKey)
	return a
}

func (*AuthCookie) CreateCookie(w http.ResponseWriter) {
	log.Println("CreateCookie")

	if encoded, err := cookieCodec.Encode("mycookie", "myvalue"); err == nil {
		cookie := http.Cookie{
			Name:  cookieName,
			Value: encoded,
			Path:  "/",
		}
		log.Println("CreateCookie ok")
		http.SetCookie(w, &cookie)
	} else {
		log.Println(err)
	}

	log.Println("end CreateCookie")
}

func (*AuthCookie) ReadCookie(r *http.Request) bool {
	log.Println("read cookie")
	if cookie, err := r.Cookie(cookieName); err == nil {
		log.Printf("Cookie %v", cookie)
		return true
	} else {
		log.Println("something wrong with the cookie")
		log.Println(err)
	}
	return false
}

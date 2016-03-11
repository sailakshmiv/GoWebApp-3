package middleware

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func LoggingHandler(next http.Handler) http.Handler {
	return handlers.CombinedLoggingHandler(os.Stdout, next)
}

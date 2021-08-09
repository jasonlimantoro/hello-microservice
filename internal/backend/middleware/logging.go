package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func WithLogging(next http.HandlerFunc) http.HandlerFunc {
	log.Out = os.Stdout
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		duration := time.Since(start)
		log.WithFields(logrus.Fields{
			"uri":      r.RequestURI,
			"method":   r.Method,
			"duration": duration,
		}).Info("Request completed")
	}
}

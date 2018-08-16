package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/leesper/holmes"
)

// LoggerMiddleware 日志中间件
func LoggerMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	start := time.Now()
	logging := fmt.Sprintf("%s -- %v %s %s %s - %s %v",
		r.RemoteAddr,
		start,
		r.Method,
		r.URL.Path,
		r.Proto,
		r.Header.Get("User-Agent"),
		time.Since(start))

	holmes.Infoln(logging)

	next(w, r)
}

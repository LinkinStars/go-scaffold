package recovery

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"

	"github.com/LinkinStars/go-scaffold/logger"
	"github.com/gin-gonic/gin"
)

// Recovery is a server middleware that recovers from any panics.
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") ||
							strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") ||
							strings.Contains(strings.ToLower(se.Error()), "protocol wrong type for socket") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Warnf("[broken pipe] [err:] %v \n [req:] %v", err, c.FullPath())
					// If the connection is dead, we can't write a status to it.
					// nolint: errcheck
					c.Abort()
					return
				}

				logger.Errorf("[Recovery from panic] [err:] %v \n [req:] %v \n [stack:] %s", err, string(httpRequest), string(debug.Stack()))
				resp := make(map[string]string)
				resp["status"] = "failure"
				resp["message"] = "InternalServerError"
				c.JSON(http.StatusInternalServerError, resp)
			}
		}()
		c.Next()
	}
}

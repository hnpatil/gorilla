package apis

import (
	"fmt"
	"math"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type responseLogWriter struct {
	gin.ResponseWriter
	ctx    *gin.Context
	logger logrus.FieldLogger

	start time.Time
}

func (rlw *responseLogWriter) Write(data []byte) (int, error) {
	rlw.writeLog(data)

	return rlw.ResponseWriter.Write(data)
}

func (rlw *responseLogWriter) writeLog(responseBody []byte) {
	path := rlw.ctx.Request.URL.Path

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	statusCode := rlw.ResponseWriter.Status()
	stop := time.Since(rlw.start)
	duration := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
	method := rlw.ctx.Request.Method
	clientIP := rlw.ctx.ClientIP()

	entry := rlw.logger.WithFields(logrus.Fields{
		"hostname":    hostname,
		"status_code": statusCode,
		"duration":    duration,
		"client_ip":   clientIP,
		"method":      method,
		"path":        path,
		"referer":     rlw.ctx.Request.Referer(),
		"data_length": len(responseBody),
		"user_agent":  rlw.ctx.Request.UserAgent(),
	})

	//  Add all possible user logger types
	identity := getUserID(rlw.ctx)
	if identity != "" {
		entry = entry.WithField("user_identity", identity)
	}

	msg := fmt.Sprintf("%s - \"%s %s\" %d (%dms)", clientIP, method, path, statusCode, duration)

	if statusCode >= http.StatusInternalServerError {
		entry.WithField("response", string(responseBody)).Error(msg)
	} else if statusCode >= http.StatusBadRequest {
		entry.WithField("response", string(responseBody)).Warn(msg)
	} else {
		entry.Info(msg)
	}
}

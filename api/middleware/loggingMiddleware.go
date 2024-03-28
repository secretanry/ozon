package middleware

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.HandlerFunc, logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		logger.Infof("Started %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		duration := time.Since(start)
		logger.Debugf("Completed %s in %s", r.URL.Path, duration)
	}
}

func UnaryServerInterceptor(logger *logrus.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		logger.Infof("Started: %s", info.FullMethod)

		resp, err := handler(ctx, req)

		logger.Debugf("Completed request: %s with error: %v", info.FullMethod, err)

		return resp, err
	}
}

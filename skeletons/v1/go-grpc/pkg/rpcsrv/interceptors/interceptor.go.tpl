package interceptors

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/selector"
	"google.golang.org/grpc"

	customlogging "{{ .App.ModuleName }}/pkg/rpcsrv/interceptors/logging"
	"{{ .App.ModuleName }}/pkg/xlog"
)

// UnaryServerInterceptors store registered grpc.UnaryServerInterceptor.
var UnaryServerInterceptors = defaultUnaryServerInterceptors()

// StreamServerInterceptors store registered grpc.StreamServerInterceptors.
var StreamServerInterceptors = defaultStreamServerInterceptors()

func defaultUnaryServerInterceptors() map[string]grpc.UnaryServerInterceptor {
	return map[string]grpc.UnaryServerInterceptor{
		"logger":   logging.UnaryServerInterceptor(customlogging.InterceptorLogger(xlog.Logger())),
		"auth":     selector.UnaryServerInterceptor(auth.UnaryServerInterceptor(grpcAuthHandler), selector.MatchFunc(allButHealthZ)),
		"recovery": recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler())),
	}
}

func defaultStreamServerInterceptors() map[string]grpc.StreamServerInterceptor {
	return map[string]grpc.StreamServerInterceptor{
		"logger":   logging.StreamServerInterceptor(customlogging.InterceptorLogger(xlog.Logger())),
		"auth":     selector.StreamServerInterceptor(auth.StreamServerInterceptor(grpcAuthHandler), selector.MatchFunc(allButHealthZ)),
		"recovery": recovery.StreamServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler())),
	}
}

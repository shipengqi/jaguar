package interceptors

import (
	"runtime/debug"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"{{ .App.ModuleName }}/pkg/xlog"
)

func grpcPanicRecoveryHandler() recovery.RecoveryHandlerFunc {
	return func(p any) (err error) {
		xlog.Errorf("recovered from panic: %v\n%s", p, debug.Stack())
		return status.Errorf(codes.Internal, "%s", p)
	}
}

package tmp

import (
	"go.temporal.io/api/serviceerror"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Grcperr(svcerr serviceerror.ServiceError) error {
	switch svcerr.Status().Code() {
	case codes.OK:
		return status.Error(codes.OK, svcerr.Status().Message())
	case codes.Canceled:
		return status.Error(codes.Canceled, svcerr.Status().Message())
	case codes.Unknown:
		return status.Error(codes.Unknown, svcerr.Status().Message())
	case codes.InvalidArgument:
		return status.Error(codes.InvalidArgument, svcerr.Status().Message())
	case codes.DeadlineExceeded:
		return status.Error(codes.DeadlineExceeded, svcerr.Status().Message())
	case codes.NotFound:
		return status.Error(codes.NotFound, svcerr.Status().Message())
	case codes.AlreadyExists:
		return status.Error(codes.AlreadyExists, svcerr.Status().Message())
	case codes.PermissionDenied:
		return status.Error(codes.PermissionDenied, svcerr.Status().Message())
	case codes.ResourceExhausted:
		return status.Error(codes.ResourceExhausted, svcerr.Status().Message())
	case codes.FailedPrecondition:
		return status.Error(codes.FailedPrecondition, svcerr.Status().Message())
	case codes.Aborted:
		return status.Error(codes.Aborted, svcerr.Status().Message())
	case codes.Unimplemented:
		return status.Error(codes.Unimplemented, svcerr.Status().Message())
	case codes.Internal:
		return status.Error(codes.Internal, svcerr.Status().Message())
	case codes.Unavailable:
		return status.Error(codes.Unavailable, svcerr.Status().Message())
	case codes.DataLoss:
		return status.Error(codes.DataLoss, svcerr.Status().Message())
	case codes.Unauthenticated:
		return status.Error(codes.Unauthenticated, svcerr.Status().Message())
	default:
		return status.Error(codes.Unknown, svcerr.Status().Message())
	}
}

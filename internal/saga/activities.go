package saga

import (
	"context"
	"time"

	"github.com/google/uuid"
	//"go.temporal.io/sdk/temporal"
	//"google.golang.org/grpc/codes"
	//"google.golang.org/grpc/status"

	"github.com/pepeunlimited/go-grpc-starter-kit/pkg/api/v1/services"
)

// Activites is the implementation of a particular task in the business logic.
// @see https://github.com/temporalio/documentation/blob/master/docs/go/activities.md

type Activities struct {
}

// @return an err or a value, err
func (a *Activities) Withdraw(ctx context.Context, request *services.CreateWithdrawRequest) (*string, error) {
	time.Sleep(20 * time.Second)
	// NOTICE: real application should request e.g. RPC/API endpoint

	//err := status.Error(codes.Internal, "intenal server error")
	// https://github.com/grpc/grpc-go/issues/2934
	//if grpcErr, ok := status.FromError(err); ok {
	//	// 	twerr
	//	return nil, temporal.NewApplicationError(grpcErr.Message(), GrpcErrorCode, grpcErr.Code())
	//} else {
	//	// unkwnerr
	//	return nil, temporal.NewApplicationError(err.Error(), "unkwnerr")
	//}
	// success return
	txId := uuid.New().String()
	return &txId, nil
}

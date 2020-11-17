package main

import (
	"context"
	"fmt"
	"os"

	pb "simple_grpc_with_downstream/internal/gen/pb/gateway"
	gateway "simple_grpc_with_downstream/internal/gen/pkg/servers/gateway"
	encoder_backend "simple_grpc_with_downstream/internal/gen/pkg/servers/gateway/encoder_backend"

	"github.com/anz-bank/pkg/log"
	"github.com/anz-bank/sysl-go/common"
	"github.com/anz-bank/sysl-go/core"
)

type AppConfig struct{}

func Encode(ctx context.Context, req *pb.EncodeRequest, client gateway.EncodeClient) (*pb.EncodeResponse, error) {
	if req.EncoderId == "rot13" {
		encoderReq := &encoder_backend.EncodingRequest{
			Content: req.Content,
		}

		encoderResponse, err := client.Encoder_backendRot13(ctx, encoderReq)
		if err != nil {
			return nil, err
		}

		return &pb.EncodeResponse{
			Content: encoderResponse.Content,
		}, nil

	} else {
		return nil, fmt.Errorf("custom response from app business logic: encoder not supported")
	}
}

func newAppServer(ctx context.Context) (core.StoppableServer, error) {
	return gateway.NewServer(ctx,
		func(ctx context.Context, cfg AppConfig) (*gateway.GrpcServiceInterface, *core.Hooks, error) {
			// FIXME auto codegen and common.MapError don't align.
			mapError := func(ctx context.Context, err error) *common.HTTPError {
				httpErr := common.MapError(ctx, err)
				return &httpErr
			}

			return &gateway.GrpcServiceInterface{
					Encode: Encode,
				}, &core.Hooks{
					MapError: mapError,
				},
				nil
		},
	)
}

func main() {
	// initialise context with pkg logger
	logger := log.NewStandardLogger()
	ctx := log.WithLogger(logger).WithConfigs(log.SetVerboseMode(true)).Onto(context.Background())

	handleError := func(err error) {
		if err != nil {
			log.Error(ctx, err)
			os.Exit(1)
		}
	}

	srv, err := newAppServer(ctx)
	handleError(err)
	err = srv.Start()
	handleError(err)
}
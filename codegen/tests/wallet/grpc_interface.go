// Code generated by sysl DO NOT EDIT.
package wallet

import (
	"context"

	pb "github.com/anz-bank/sysl-go/codegen/tests/cardspb"
)

// Apple Client
type AppleClient struct {
}

// GrpcServiceInterface for Wallet
type GrpcServiceInterface struct {
	Apple func(ctx context.Context, req *pb.AppleRequest, client AppleClient) (*pb.AppleResponse, error)
}
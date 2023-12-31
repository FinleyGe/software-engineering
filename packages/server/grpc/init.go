package grpc

import (
	"log"
	"net"
	. "se/config"
	grpcapi "se/grpc/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type VitalSignServer struct {
	grpcapi.UnimplementedVitalSignServiceServer
}

type BellServer struct {
	grpcapi.UnimplementedBellServiceServer
}

var Lis net.Listener
var GrpcServer *grpc.Server

func init() {
	var err error
	Lis, err = net.Listen("tcp", "0.0.0.0:"+Config.Grpc.Port)
	if err != nil {
		log.Fatalf("[GRPC] failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	GrpcServer = grpc.NewServer(opts...)
	reflection.Register(GrpcServer)
	grpcapi.RegisterVitalSignServiceServer(GrpcServer, &VitalSignServer{})
	grpcapi.RegisterBellServiceServer(GrpcServer, &BellServer{})
}

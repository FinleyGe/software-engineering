package device

import (
	"context"
	grpcapi "hardware/grpc/api"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Bell struct {
	DepartmentID uint64 `json:"department_id"`
	BedID        uint64 `json:"bed_id"`
	IsRunning    bool   `json:"is_running"`
}

func (b *Bell) Run() {
	b.IsRunning = true
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(os.Getenv("SE_GRPC_ADDRESS"), opts...)
	log.Println(os.Getenv("SE_GRPC_ADDRESS"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer conn.Close()
	client := grpcapi.NewBellServiceClient(conn)

	stream, err := client.BellStream(
		context.Background(),
		&grpcapi.BellRequest{
			DepartmentID: b.DepartmentID,
		},
	)
	for {
		if !b.IsRunning {
			break
		}
		r, err := stream.Recv()
		if err != nil {
			log.Println("BellStream error:", err)
			break
		}
		log.Println("BellStream received:", r.BedID)
	}
}

func (b *Bell) Stop() {
	b.IsRunning = false
}

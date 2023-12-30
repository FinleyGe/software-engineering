package grpc

import (
	"context"
	grpcapi "hardware/grpc/api"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Run() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(os.Getenv("SE_GRPC_ADDRESS"), opts...)
	log.Println(os.Getenv("SE_GRPC_ADDRESS"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer conn.Close()

	client := grpcapi.NewVitalSignServiceClient(conn)
	stream, err := client.PushVitalSign(context.Background())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer stream.CloseSend()
	for {
		err = stream.Send(&grpcapi.VitalSign{
			ID:            1,
			HeartRate:     60,
			Temperature:   36.5,
			BloodPressure: "120/80",
			BloodOxygen:   98,
			Time:          "2020-01-01 00:00:00",
			BreathingRate: 20,
			Sense:         "good",
		})
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		time.Sleep(1 * time.Second)
		log.Println("send")
	}
}

package device

import (
	"context"
	"fmt"
	grpcapi "hardware/grpc/api"
	"log"
	"math/rand"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Monitor struct {
	BedID     uint64 `json:"bed_id"`
	IsRunning bool   `json:"is_running"`
}

func (m *Monitor) Run() {
	m.IsRunning = true
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
		var heartRate uint64 = rand.Uint64()%150 + 60
		var temperature float64 = rand.Float64()*4 + 35
		var bloodPressure string = fmt.Sprintf("%v/%v", rand.Uint64()%100+60, rand.Uint64()%100+60)
		var bloodOxygen float64 = rand.Float64()*2 + 96
		var breathingRate uint64 = rand.Uint64()%20 + 10
		var sense string = "normal"
		if heartRate > 160 || heartRate < 60 {
			sense = "non-sense"
		}
		if temperature > 39 || temperature < 35 {
			sense = "non-sense"
		}
		if bloodOxygen < 90 {
			sense = "non-sense"
		}
		err = stream.Send(&grpcapi.VitalSign{
			ID:            m.BedID,
			HeartRate:     heartRate,
			Temperature:   temperature,
			BloodPressure: bloodPressure,
			BloodOxygen:   bloodOxygen,
			Time:          timestamppb.Now(),
			BreathingRate: breathingRate,
			Sense:         sense,
		})
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		time.Sleep(1 * time.Second)
		log.Println(fmt.Sprintf("Device %v: sended!", m.BedID))
		if !m.IsRunning {
			log.Println(fmt.Sprintf("Device %v: stopped!", m.BedID))
			break
		}
	}
}

func (m *Monitor) Call() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(os.Getenv("SE_GRPC_ADDRESS"), opts...)
	log.Println(os.Getenv("SE_GRPC_ADDRESS"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer conn.Close()
	client := grpcapi.NewVitalSignServiceClient(conn)
	client.Call(context.Background(),
		&grpcapi.CallRequest{
			BedID: m.BedID,
		},
	)
}

func (m *Monitor) Stop() {
	m.IsRunning = false
}

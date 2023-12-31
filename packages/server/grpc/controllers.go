package grpc

import (
	"context"
	"errors"
	"log"
	"se/db/model"
	"se/grpc/api"
	"se/utility"
	"sync"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (s *VitalSignServer) PushVitalSign(
	stream grpcapi.VitalSignService_PushVitalSignServer,
) error {
	for {
		vitalSign, err := stream.Recv()
		if err != nil {
			log.Println(err)
			break
		}
		l := model.VitalSign{
			BedID:         vitalSign.ID,
			Time:          utility.Time(vitalSign.Time.AsTime()),
			HeartRate:     vitalSign.HeartRate,
			BloodPressure: vitalSign.BloodPressure,
			Temperature:   vitalSign.Temperature,
			Sense:         vitalSign.Sense,
			BloodOxygen:   vitalSign.BloodOxygen,
			BreathingRate: vitalSign.BreathingRate,
		}
		l.Add()
	}
	return nil
}

type callData struct {
	DepartmentID uint64
	BedID        uint64
}

var callChan = make(chan callData)
var availableDepartment = sync.Map{}

func (s *BellServer) BellStream(
	r *grpcapi.BellRequest,
	stream grpcapi.BellService_BellStreamServer,
) error {
	availableDepartment.Store(r.DepartmentID, true)
	log.Println("get a call request, department id:", r.DepartmentID)
	defer availableDepartment.Delete(r.DepartmentID)
	for {
		call := <-callChan
		if call.DepartmentID == r.DepartmentID {
			stream.Send(&grpcapi.BellResponse{
				BedID: call.BedID,
			})
		}
	}
}

func (s *BellServer) Call(
	c context.Context,
	r *grpcapi.CallRequest,
) (emptypb.Empty, error) {
	patient := model.Patient{}
	patient.ID = r.BedID
	departmentID := patient.GetDepartmentID()
	if _, ok := availableDepartment.Load(departmentID); !ok {
		return emptypb.Empty{}, errors.New("department not available")
	}
	callChan <- callData{
		DepartmentID: departmentID,
		BedID:        r.BedID,
	}
	return emptypb.Empty{}, nil
}

func Call(patient_id uint64) error {
	patient := model.Patient{}
	patient.ID = patient_id
	departmentID := patient.GetDepartmentID()
	log.Println("call department id:", departmentID)
	if _, ok := availableDepartment.Load(departmentID); !ok {
		return errors.New("department not available")
	}
	callChan <- callData{
		DepartmentID: departmentID,
		BedID:        patient_id,
	}
	return nil
}

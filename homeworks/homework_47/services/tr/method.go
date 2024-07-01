package tr

import (
	"context"
	pb "homeworks/homework_47/genproto/transport"
)

func (t *TransportRepo) GetBusSchedule(ctx context.Context, in *pb.BusNumber) (*pb.BusSchudle, error) {
	location := []*pb.Location{
		{
			Location: "any where",
			Time:     "every 27 minutes",
		},
	}
	return &pb.BusSchudle{BusNumber: in.Bus, Stations: location}, nil
}

func (t *TransportRepo) TrucBusLocation(ctx context.Context, in *pb.Location) (*pb.BusWithLocations, error) {
	bus := []*pb.Bus{
		{
			Number: 45,
			Time:   "every 22 minutes",
		},
	}
	return &pb.BusWithLocations{Busses: bus}, nil
}

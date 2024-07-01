package wth

import (
	"context"
	pb "homeworks/homework_47/genproto/weather"
)

func (w *WeatherRepo) GetCurrentWeather(ctx context.Context, in *pb.Void) (*pb.WeatherDaily, error) {
	weather := pb.WeatherCondition{Temperature: 35, Humidity: "34%", WindSpeed: 11, Condition: "sunny"}
	return &pb.WeatherDaily{Date: "2024-06-28", Location: "Tashkent", Weather: &weather}, nil
}

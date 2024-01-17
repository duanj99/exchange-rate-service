package handler

import (
	"CurrencyExchangeService/controller"
	"CurrencyExchangeService/logger"
	"CurrencyExchangeService/repository"
	server_grpc "CurrencyExchangeService/server_grpc/protogen"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
	"time"
	"unsafe"
)

type ExchangeRateGRPCServer struct {
	server_grpc.UnimplementedExchangeRateServiceServer // interface
	Logger                                             *logger.ServiceLogger
	Controller                                         *controller.ExchangeRateController
}

func (e ExchangeRateGRPCServer) GetLatestRate(ctx context.Context, request *server_grpc.GetLatestRateRequest) (*server_grpc.GetLatestRateResponse, error) {
	//return nil, status.Errorf(codes.Unimplemented, "method GetLatestRate not implemented")
	e.Logger.Info("Someone is calling the grpc GetLatestRate endpoint")
	resp := e.Controller.GetLatestRate()

	respTime := resp.InsertTimeStamp.Time().Unix()

	grpcResp := &server_grpc.GetLatestRateResponse{
		Rate: &server_grpc.ExchangeRate{
			Rates:        resp.Rates,
			BaseCurrency: resp.BaseCurrency,
			Time: &timestamppb.Timestamp{
				Seconds: respTime,
			},
		},
	}

	return grpcResp, nil
}

func (e ExchangeRateGRPCServer) GetRangeRate(ctx context.Context, request *server_grpc.GetRangeRateRequest) (*server_grpc.GetRangeRateResponse, error) {
	//TODO implement me
	//return nil, status.Errorf(codes.Unimplemented, "method GetLatestRate not implemented")
	e.Logger.Info("Someone is calling the grpc GetRange endpoint")
	startTime, _ := time.Parse("2006-01-02", request.GetStartDate())
	endTime, _ := time.Parse("2006-01-02", request.GetEndDate())
	input := repository.RangeRateRequest{
		StartTime: primitive.DateTime(startTime.UnixMilli()),
		EndTime:   primitive.DateTime(endTime.UnixMilli()),
	}

	e.Logger.Info(fmt.Sprintf("%+v", input))

	respStruct := e.Controller.GetRangeRates(input)

	var rateList []*server_grpc.ExchangeRate

	for _, resp := range respStruct {
		rate := &server_grpc.ExchangeRate{
			Rates:        resp.Rates,
			BaseCurrency: resp.BaseCurrency,
			Time: &timestamppb.Timestamp{
				Seconds: resp.InsertTimeStamp.Time().Unix(),
			},
		}
		rateList = append(rateList, rate)
	}

	respGrpc := &server_grpc.GetRangeRateResponse{Rates: rateList}
	e.Logger.Info(fmt.Sprintf("the respGrpc Data: %s, Size:, %d", respGrpc.ProtoReflect(), unsafe.Sizeof(respGrpc)))

	return respGrpc, nil

}

// InitGRPCHandler Initialize the gRPC Handler
func InitGRPCHandler(
	svcLogger *logger.ServiceLogger,
	svcController *controller.ExchangeRateController,
) {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	serviceRegistrar := grpc.NewServer()
	service := &ExchangeRateGRPCServer{
		Logger:     svcLogger,
		Controller: svcController,
	}
	server_grpc.RegisterExchangeRateServiceServer(serviceRegistrar, service)
	reflection.Register(serviceRegistrar)
	err = serviceRegistrar.Serve(lis)

	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}

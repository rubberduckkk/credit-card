package app

import (
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	pb "github.com/rubberduckkk/credit-card/api/pb/credit-card"
)

func RunGRPCServer(listener net.Listener) error {
	grpcServer := grpc.NewServer()
	pb.RegisterCreditCardServer(grpcServer, GetCreditCardServer())

	logrus.Info("starting grpc server ...")
	return grpcServer.Serve(listener)
}

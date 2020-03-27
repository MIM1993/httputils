package server

import (
	"context"
	realtypb "gitlab.bj.sensetime.com/SenseRealEstate/Schema/realty.git"
	"net"
	"net/http"
	"github.com/sirupsen/logrus"
	gwrt "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

// grpc + restful  样例
func startRpc() error {

	config := utils.Global.Config
	grpcEndpoint := config.GrpcEndpoint
	httpEndpoint := config.ListenAddressHTTP

	go func() {
		grpcListener, err := net.Listen("tcp", grpcEndpoint)
		if err != nil {
			logrus.Fatalf("failed to listen: %v", err)
		}
		grpcSvr := grpc.NewServer()
		realtypb.RegisterRealtyServer(grpcSvr, service.NewGrpcService())  //注册
		logrus.Infof("grpc Server addr: %v", grpcEndpoint)
		err = grpcSvr.Serve(grpcListener)
		if err != nil {
			logrus.Fatalf("grpcSvr.Server err:%v", err)
		}
	}()

	// restful
	go func() {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		mux := gwrt.NewServeMux()
		opts := []grpc.DialOption{grpc.WithInsecure()}
		err := realtypb.RegisterRealtyHandlerFromEndpoint(ctx, mux, grpcEndpoint, opts)  //注册
		if err != nil {
			logrus.Fatalf("pb.RegisterEdgeServiceHandlerFromEndpoint err:%v", err)
		}

		logrus.Infof("http.ListenAndServe addr: %v", httpEndpoint)
		err = http.ListenAndServe(httpEndpoint, mux)
		if err != nil {
			logrus.Fatalf("http.ListenAndServe err:%v", err)
		}
	}()

	return nil
}

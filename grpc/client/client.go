package client

import (
	"github.com/sercand/kuberesolver"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
	"strings"

	realtypb "gitlab.bj.sensetime.com/SenseRealEstate/Schema/realty.git"
)

// "kubernetes:///progress-service-svc.just-4-test:50051"
// "10.152.26.147:50051"
func newConn(address string) (*grpc.ClientConn, error) {
	if !strings.Contains(address, "kubernetes:") {
		return grpc.Dial(address, grpc.WithInsecure())
	}

	client, err := kuberesolver.NewInClusterK8sClient()
	if err != nil {
		return nil, err
	}
	resolver.Register(kuberesolver.NewBuilder(client, "kubernetes"))
	return grpc.Dial(address,
		grpc.WithInsecure(),
		grpc.WithBalancerName(roundrobin.Name),
	)
}

// TODO 复用，负载均衡
func RealtyClient() (realtypb.RealtyClient, error) {

	conn, err := newConn(Global.Config.GrpcServerAddress)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}

	// defer conn.Close()
	return realtypb.NewRealtyClient(conn), nil
}

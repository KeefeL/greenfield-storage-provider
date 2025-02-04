package gfspapp

import (
	"context"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"

	"github.com/bnb-chain/greenfield-storage-provider/base/types/gfspserver"
	"github.com/bnb-chain/greenfield-storage-provider/pkg/log"
	utilgrpc "github.com/bnb-chain/greenfield-storage-provider/util/grpc"
)

const (
	// MaxServerCallMsgSize defines the max message size for grpc server
	MaxServerCallMsgSize = 3 * 1024 * 1024 * 1024
)

func DefaultGrpcServerOptions() []grpc.ServerOption {
	var options []grpc.ServerOption
	options = append(options, grpc.MaxRecvMsgSize(MaxServerCallMsgSize))
	options = append(options, grpc.MaxSendMsgSize(MaxServerCallMsgSize))
	return options
}

func (g *GfSpBaseApp) newRpcServer(options ...grpc.ServerOption) {
	options = append(options, DefaultGrpcServerOptions()...)
	options = append(options, GeKeepAliveServerOptions()...)
	if g.EnableMetrics() {
		options = append(options, utilgrpc.GetDefaultServerInterceptor()...)
	}
	g.server = grpc.NewServer(options...)
	gfspserver.RegisterGfSpApprovalServiceServer(g.server, g)
	gfspserver.RegisterGfSpAuthenticationServiceServer(g.server, g)
	gfspserver.RegisterGfSpDownloadServiceServer(g.server, g)
	gfspserver.RegisterGfSpManageServiceServer(g.server, g)
	gfspserver.RegisterGfSpP2PServiceServer(g.server, g)
	gfspserver.RegisterGfSpResourceServiceServer(g.server, g)
	gfspserver.RegisterGfSpReceiveServiceServer(g.server, g)
	gfspserver.RegisterGfSpSignServiceServer(g.server, g)
	gfspserver.RegisterGfSpUploadServiceServer(g.server, g)
	gfspserver.RegisterGfSpQueryTaskServiceServer(g.server, g)
	reflection.Register(g.server)
}

func (g *GfSpBaseApp) StartRPCServer(ctx context.Context) error {
	lis, err := net.Listen("tcp", g.grpcAddress)
	if err != nil {
		log.Errorw("failed to listen tcp address", "address", g.grpcAddress, "error", err)
		return err
	}
	go func() {
		if err = g.server.Serve(lis); err != nil {
			log.Errorw("failed to start gf-sp app grpc server", "error", err)
		}
	}()
	return nil
}

func (g *GfSpBaseApp) StopRPCServer(ctx context.Context) error {
	g.server.GracefulStop()
	return nil
}

func GetRPCRemoteAddress(ctx context.Context) string {
	var addr string
	if pr, ok := peer.FromContext(ctx); ok {
		if tcpAddr, ok := pr.Addr.(*net.TCPAddr); ok {
			addr = tcpAddr.IP.String()
		} else {
			addr = pr.Addr.String()
		}
	}
	return addr
}

// GeKeepAliveServerOptions returns keepalive gRPC server options
func GeKeepAliveServerOptions() []grpc.ServerOption {
	var kaep = keepalive.EnforcementPolicy{
		MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
		PermitWithoutStream: true,            // Allow pings even when there are no active streams
	}

	var kasp = keepalive.ServerParameters{
		MaxConnectionIdle:     10 * time.Second, // If a client is idle for 10 seconds, send a GOAWAY
		MaxConnectionAge:      10 * time.Second, // If any connection is alive for more than 10 seconds, send a GOAWAY
		MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
		Time:                  5 * time.Second,  // Ping the client if it is idle for 5 seconds to ensure the connection is still active
		Timeout:               1 * time.Second,  // Wait 1 second for the ping ack before assuming the connection is dead
	}

	return append([]grpc.ServerOption{}, grpc.KeepaliveEnforcementPolicy(kaep), grpc.KeepaliveParams(kasp))
}

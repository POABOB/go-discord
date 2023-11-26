package main

import (
	"flag"
	"fmt"

	"github.com/POABOB/go-discord/apps/servers/rpc/internal/config"
	memberrpcServer "github.com/POABOB/go-discord/apps/servers/rpc/internal/server/memberrpc"
	serverrpcServer "github.com/POABOB/go-discord/apps/servers/rpc/internal/server/serverrpc"
	"github.com/POABOB/go-discord/apps/servers/rpc/internal/svc"
	"github.com/POABOB/go-discord/apps/servers/rpc/pb/rpc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/servers.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		rpc.RegisterServerRpcServer(grpcServer, serverrpcServer.NewServerRpcServer(ctx))
		rpc.RegisterMemberRpcServer(grpcServer, memberrpcServer.NewMemberRpcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

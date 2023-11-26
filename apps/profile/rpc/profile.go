package main

import (
	"flag"
	"fmt"

	"github.com/POABOB/go-discord/apps/profile/rpc/internal/config"
	rpcServer "github.com/POABOB/go-discord/apps/profile/rpc/internal/server/rpc"
	"github.com/POABOB/go-discord/apps/profile/rpc/internal/svc"
	"github.com/POABOB/go-discord/apps/profile/rpc/pb/rpc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/profile.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		rpc.RegisterRpcServer(grpcServer, rpcServer.NewRpcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc servers at %s...\n", c.ListenOn)
	s.Start()
}

package svc

import (
	"github.com/POABOB/go-discord/apps/app/api/internal/config"
	profile "github.com/POABOB/go-discord/apps/profile/rpc/client/rpc"
	member "github.com/POABOB/go-discord/apps/servers/rpc/client/memberrpc"
	server "github.com/POABOB/go-discord/apps/servers/rpc/client/serverrpc"
	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	Profile profile.Rpc
	Server  server.ServerRpc
	Member  member.MemberRpc
	Clerk   clerk.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	client, _ := clerk.NewClient("sk_test_WgMRR4t2EztZwXZRfYPW85QjKFvAsh78DR1QEJpT32")
	return &ServiceContext{
		Config:  c,
		Profile: profile.NewRpc(zrpc.MustNewClient(c.ProfileRpc)),
		Server:  server.NewServerRpc(zrpc.MustNewClient(c.ServerRpc)),
		Member:  member.NewMemberRpc(zrpc.MustNewClient(c.ServerRpc)),
		Clerk:   client,
	}
}

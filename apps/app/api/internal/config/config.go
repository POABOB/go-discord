package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	// jwt 需要的密鑰和過期時間
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	ProfileRpc zrpc.RpcClientConf
	ServerRpc  zrpc.RpcClientConf
	MemberRpc  zrpc.RpcClientConf
}

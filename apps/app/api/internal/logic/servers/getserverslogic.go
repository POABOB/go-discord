package servers

import (
	"context"
	serverRPC "github.com/POABOB/go-discord/apps/servers/rpc/pb/rpc"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/status"

	"github.com/POABOB/go-discord/apps/app/api/internal/svc"
	"github.com/POABOB/go-discord/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetServersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetServersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetServersLogic {
	return &GetServersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetServersLogic) GetServers(req *types.GetServersReq) ([]types.GetServerRes, error) {
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PagePerNum == 0 || req.PagePerNum > 1000 {
		req.PagePerNum = 10
	}

	var rpcReq serverRPC.GetServersReq
	if err := copier.Copy(&rpcReq, &req); err != nil {
		return nil, status.Error(500, err.Error())
	}

	temp, err := l.svcCtx.Server.GetServers(l.ctx, &rpcReq)
	if err != nil {
		return nil, status.Error(400, err.Error())
	}

	resp := make([]types.GetServerRes, 0)
	if err = copier.Copy(&resp, &temp.Servers); err != nil {
		return nil, status.Error(500, err.Error())
	}
	return resp, nil
}

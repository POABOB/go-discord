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

type GetUniqueServerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUniqueServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUniqueServerLogic {
	return &GetUniqueServerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUniqueServerLogic) GetUniqueServer(req *types.GetServerReq) (*types.GetServerRes, error) {
	temp, err := l.svcCtx.Server.GetUniqueServer(l.ctx, &serverRPC.GetServerReq{
		Id:        req.ServerId,
		ProfileId: l.ctx.Value("id").(string),
	})
	if err != nil {
		return nil, status.Error(400, err.Error())
	}

	resp := &types.GetServerRes{}
	if err = copier.Copy(&resp, &temp.Server); err != nil {
		return nil, status.Error(500, err.Error())
	}
	if err = copier.Copy(&resp.Members, &temp.Relation.Members); err != nil {
		return nil, status.Error(500, err.Error())
	}
	if err = copier.Copy(&resp.Channels, &temp.Relation.Channels); err != nil {
		return nil, status.Error(500, err.Error())
	}
	return resp, nil
}

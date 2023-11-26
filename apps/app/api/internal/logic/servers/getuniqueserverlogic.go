package servers

import (
	"context"
	serverRPC "github.com/POABOB/go-discord/apps/servers/rpc/pb/rpc"
	"github.com/jinzhu/copier"

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

func (l *GetUniqueServerLogic) GetUniqueServer(req *types.GetServerReq) (resp *types.GetServerRes, err error) {
	var temp *serverRPC.GetServerRes
	temp, err = l.svcCtx.Server.GetUniqueServer(l.ctx, &serverRPC.GetServerReq{
		Id:        req.ServerId,
		ProfileId: l.ctx.Value("id").(string),
	})
	if err != nil {
		return
	}
	resp = new(types.GetServerRes)
	if err = copier.Copy(&resp, &temp.Server); err != nil {
		return
	}
	if err = copier.Copy(&resp.Members, &temp.Relation.Members); err != nil {
		return
	}
	if err = copier.Copy(&resp.Channels, &temp.Relation.Channels); err != nil {
		return
	}
	//fmt.Println(temp.Relation.Members)
	return
}

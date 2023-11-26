package member

import (
	"context"
	serverRPC "github.com/POABOB/go-discord/apps/servers/rpc/pb/rpc"
	"github.com/jinzhu/copier"

	"github.com/POABOB/go-discord/apps/app/api/internal/svc"
	"github.com/POABOB/go-discord/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMemberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMemberLogic {
	return &DeleteMemberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMemberLogic) DeleteMember(req *types.DeleteMemberReq) (resp *types.GetServerRes, err error) {
	var temp *serverRPC.GetServerRes
	profileId := l.ctx.Value("id").(string)
	// Call RPC
	if temp, err = l.svcCtx.Member.DeleteMember(l.ctx, &serverRPC.DeleteMemberReq{
		Id:        req.Id,
		ServerId:  req.ServerId,
		ProfileId: profileId,
	}); err != nil {
		return
	}

	resp = &types.GetServerRes{}
	if err = copier.Copy(&resp, &temp.Server); err != nil {
		return
	}
	if err = copier.Copy(&resp.Members, &temp.Relation.Members); err != nil {
		return
	}
	if err = copier.Copy(&resp.Channels, &temp.Relation.Channels); err != nil {
		return
	}
	return
}

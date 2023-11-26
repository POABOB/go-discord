package member

import (
	"context"
	serverRPC "github.com/POABOB/go-discord/apps/servers/rpc/pb/rpc"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/status"

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

func (l *DeleteMemberLogic) DeleteMember(req *types.DeleteMemberReq) (*types.GetServerRes, error) {
	profileId := l.ctx.Value("id").(string)
	// Call RPC
	temp, err := l.svcCtx.Member.DeleteMember(l.ctx, &serverRPC.DeleteMemberReq{
		Id:        req.Id,
		ServerId:  req.ServerId,
		ProfileId: profileId,
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

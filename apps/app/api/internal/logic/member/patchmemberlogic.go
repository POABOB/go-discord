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

type PatchMemberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPatchMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PatchMemberLogic {
	return &PatchMemberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PatchMemberLogic) PatchMember(req *types.PatchMemberReq) (*types.GetServerRes, error) {
	temp, err := l.svcCtx.Member.PatchMember(l.ctx, &serverRPC.PatchMemberReq{
		Id:        req.Id,
		ServerId:  req.ServerId,
		ProfileId: req.ProfileId,
		Role:      req.Role,
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

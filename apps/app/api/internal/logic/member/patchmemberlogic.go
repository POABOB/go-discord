package member

import (
	"context"
	serverRPC "github.com/POABOB/go-discord/apps/servers/rpc/pb/rpc"
	"github.com/jinzhu/copier"

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

func (l *PatchMemberLogic) PatchMember(req *types.PatchMemberReq) (resp *types.GetServerRes, err error) {
	var temp *serverRPC.GetServerRes

	temp, err = l.svcCtx.Member.PatchMember(l.ctx, &serverRPC.PatchMemberReq{
		Id:        req.Id,
		ServerId:  req.ServerId,
		ProfileId: req.ProfileId,
		Role:      req.Role,
	})
	if err != nil {
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
	//fmt.Println(temp.Relation.Members)
	return
}

package member

import (
	"context"
	serverRPC "github.com/POABOB/go-discord/apps/servers/rpc/pb/rpc"
	"github.com/jinzhu/copier"

	"github.com/POABOB/go-discord/apps/app/api/internal/svc"
	"github.com/POABOB/go-discord/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InviteMemberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInviteMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InviteMemberLogic {
	return &InviteMemberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InviteMemberLogic) InviteMember(req *types.InviteMemberReq) (resp *types.GetServerRes, err error) {
	var temp *serverRPC.GetServerRes

	temp, err = l.svcCtx.Member.InviteMember(l.ctx, &serverRPC.InviteMemberReq{
		ProfileId:  l.ctx.Value("id").(string),
		InviteCode: req.InviteCode,
	})
	if err != nil {
		return
	}
	resp = &types.GetServerRes{}
	if err = copier.Copy(&resp, &temp.Server); err != nil {
		return
	}
	return
}

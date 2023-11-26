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

func (l *InviteMemberLogic) InviteMember(req *types.InviteMemberReq) (*types.GetServerRes, error) {
	temp, err := l.svcCtx.Member.InviteMember(l.ctx, &serverRPC.InviteMemberReq{
		ProfileId:  l.ctx.Value("id").(string),
		InviteCode: req.InviteCode,
	})
	if err != nil {
		return nil, status.Error(400, err.Error())
	}

	resp := &types.GetServerRes{}
	if err = copier.Copy(&resp, &temp.Server); err != nil {
		return nil, status.Error(500, err.Error())
	}
	return resp, nil
}

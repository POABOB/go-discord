package servers

import (
	"context"
	serverRPC "github.com/POABOB/go-discord/apps/servers/rpc/pb/rpc"

	"github.com/POABOB/go-discord/apps/app/api/internal/svc"
	"github.com/POABOB/go-discord/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PatchServerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPatchServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PatchServerLogic {
	return &PatchServerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PatchServerLogic) PatchServer(req *types.PatchServerReq) error {
	// Call RPC
	if _, err := l.svcCtx.Server.PatchServer(l.ctx, &serverRPC.PatchServerReq{
		Id:         req.Id,
		Name:       req.Name,
		ImageUrl:   req.ImageUrl,
		ProfileId:  req.ProfileId,
		InviteCode: req.InviteCode,
	}); err != nil {
		return err
	}

	return nil
}

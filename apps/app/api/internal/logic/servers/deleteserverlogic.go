package servers

import (
	"context"
	serverRPC "github.com/POABOB/go-discord/apps/servers/rpc/pb/rpc"

	"github.com/POABOB/go-discord/apps/app/api/internal/svc"
	"github.com/POABOB/go-discord/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteServerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteServerLogic {
	return &DeleteServerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteServerLogic) DeleteServer(req *types.DeleteServerReq) error {
	profileId := l.ctx.Value("id").(string)

	// Call RPC
	if _, err := l.svcCtx.Server.DeleteServer(l.ctx, &serverRPC.DeleteServerReq{
		Id:        req.Id,
		ProfileId: profileId,
	}); err != nil {
		return err
	}

	return nil
}

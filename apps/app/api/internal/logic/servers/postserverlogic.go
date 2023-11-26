package servers

import (
	"context"
	serverRPC "github.com/POABOB/go-discord/apps/servers/rpc/pb/rpc"
	"google.golang.org/grpc/status"

	"github.com/POABOB/go-discord/apps/app/api/internal/svc"
	"github.com/POABOB/go-discord/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostServerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostServerLogic {
	return &PostServerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostServerLogic) PostServer(req *types.PostServerReq) error {
	profileId := l.ctx.Value("id").(string)

	// Call RPC
	if _, err := l.svcCtx.Server.PostServer(l.ctx, &serverRPC.PostServerReq{
		Name:      req.Name,
		ImageUrl:  req.ImageUrl,
		ProfileId: profileId,
	}); err != nil {
		return status.Error(400, err.Error())
	}

	return nil
}

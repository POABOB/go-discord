package serverrpclogic

import (
	"context"
	"github.com/POABOB/go-discord/prisma/db"

	"github.com/POABOB/go-discord/apps/servers/rpc/internal/svc"
	"github.com/POABOB/go-discord/apps/servers/rpc/pb/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteServerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteServerLogic {
	return &DeleteServerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteServerLogic) DeleteServer(in *rpc.DeleteServerReq) (*rpc.Empty, error) {
	ctx := context.Background()

	// delete
	_, err := l.svcCtx.PrismaClient.Server.FindMany(
		db.Server.ID.Equals(in.Id),
		db.Server.ProfileID.Equals(in.ProfileId),
	).Delete().Exec(ctx)
	if err != nil {
		return &rpc.Empty{}, err
	}

	return &rpc.Empty{}, nil
}

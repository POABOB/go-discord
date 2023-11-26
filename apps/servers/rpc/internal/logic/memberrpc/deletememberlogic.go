package memberrpclogic

import (
	"context"
	"github.com/POABOB/go-discord/apps/servers/rpc/internal/logic/serverrpc"
	"github.com/POABOB/go-discord/prisma/db"

	"github.com/POABOB/go-discord/apps/servers/rpc/internal/svc"
	"github.com/POABOB/go-discord/apps/servers/rpc/pb/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMemberLogic {
	return &DeleteMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMemberLogic) DeleteMember(in *rpc.DeleteMemberReq) (*rpc.GetServerRes, error) {
	ctx := context.Background()

	deleteMember := l.svcCtx.PrismaClient.Member.
		FindMany(
			db.Member.ID.Equals(in.Id),
			db.Member.ServerID.Equals(in.ServerId),
			db.Member.ProfileID.Not(in.ProfileId),
		).
		Delete().Tx()
	if err := l.svcCtx.PrismaClient.Prisma.Transaction(deleteMember).Exec(ctx); err != nil {
		return nil, err
	}

	getUniqueServer := serverrpclogic.NewGetUniqueServerLogic(l.ctx, l.svcCtx)
	return getUniqueServer.GetUniqueServer(&rpc.GetServerReq{Id: in.ServerId, ProfileId: in.ProfileId})
}

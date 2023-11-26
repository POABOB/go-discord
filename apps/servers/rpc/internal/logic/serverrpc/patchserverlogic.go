package serverrpclogic

import (
	"context"
	"github.com/POABOB/go-discord/prisma/db"

	"github.com/POABOB/go-discord/apps/servers/rpc/internal/svc"
	"github.com/POABOB/go-discord/apps/servers/rpc/pb/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PatchServerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPatchServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PatchServerLogic {
	return &PatchServerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PatchServerLogic) PatchServer(in *rpc.PatchServerReq) (*rpc.EmptyRes, error) {
	ctx := context.Background()

	// patch
	_, err := l.svcCtx.PrismaClient.Server.
		FindMany(
			db.Server.ID.Equals(in.Id),
			db.Server.ProfileID.Equals(in.ProfileId),
		).
		Update(
			db.Server.Name.Set(in.Name),
			db.Server.ImageURL.Set(in.ImageUrl),
			db.Server.InviteCode.Set(in.InviteCode),
		).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

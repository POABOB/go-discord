package rpclogic

import (
	"context"
	"github.com/POABOB/go-discord/prisma/db"

	"github.com/POABOB/go-discord/apps/profile/rpc/internal/svc"
	"github.com/POABOB/go-discord/apps/profile/rpc/pb/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProfileLogic {
	return &DeleteProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteProfileLogic) DeleteProfile(in *rpc.DeleteProfileReq) (*rpc.Empty, error) {
	ctx := context.Background()

	// delete
	_, err := l.svcCtx.PrismaClient.Profile.FindUnique(db.Profile.ID.Equals(in.Id)).Delete().Exec(ctx)
	if err != nil {
		return &rpc.Empty{}, err
	}

	return &rpc.Empty{}, nil
}

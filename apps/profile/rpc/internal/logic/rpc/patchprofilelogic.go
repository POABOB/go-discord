package rpclogic

import (
	"context"
	"github.com/POABOB/go-discord/prisma/db"

	"github.com/POABOB/go-discord/apps/profile/rpc/internal/svc"
	"github.com/POABOB/go-discord/apps/profile/rpc/pb/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PatchProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPatchProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PatchProfileLogic {
	return &PatchProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PatchProfileLogic) PatchProfile(in *rpc.PatchProfileReq) (*rpc.EmptyRes, error) {
	ctx := context.Background()

	// patch
	_, err := l.svcCtx.PrismaClient.Profile.
		FindUnique(db.Profile.ID.Equals(in.Id)).
		Update(
			db.Profile.Name.Set(in.Name),
			db.Profile.Email.Set(in.Email),
			db.Profile.ImageURL.Set(in.ImageUrl),
		).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

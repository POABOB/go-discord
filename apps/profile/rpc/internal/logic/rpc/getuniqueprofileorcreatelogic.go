package rpclogic

import (
	"context"
	"github.com/POABOB/go-discord/apps/profile/rpc/internal/svc"
	"github.com/POABOB/go-discord/apps/profile/rpc/pb/rpc"
	"github.com/POABOB/go-discord/prisma/db"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUniqueProfileOrCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUniqueProfileOrCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUniqueProfileOrCreateLogic {
	return &GetUniqueProfileOrCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUniqueProfileOrCreateLogic) GetUniqueProfileOrCreate(in *rpc.GetUniqueProfileOrCreateReq) (*rpc.GetUniqueProfileOrCreateRes, error) {
	ctx := context.Background()

	// find a single post
	profileModel, err := l.svcCtx.PrismaClient.Profile.FindUnique(
		db.Profile.UserID.Equals(in.UserId),
	).Exec(ctx)
	if err != nil && err.Error() == "ErrNotFound" {
		// create a profile
		if profileModel, err = l.svcCtx.PrismaClient.Profile.CreateOne(
			db.Profile.UserID.Set(in.UserId),
			db.Profile.Name.Set(in.Name),
			db.Profile.ImageURL.Set(in.ImageUrl),
			db.Profile.Email.Set(in.Email),
		).Exec(ctx); err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	profile := rpc.GetUniqueProfileOrCreateRes{
		Id:       profileModel.ID,
		UserId:   profileModel.UserID,
		Name:     profileModel.Name,
		ImageUrl: profileModel.ImageURL,
		Email:    profileModel.Email,
	}

	return &profile, nil
}

package profile

import (
	"context"
	"github.com/POABOB/go-discord/apps/app/api/internal/svc"
	"github.com/POABOB/go-discord/apps/app/api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUniqueProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUniqueProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUniqueProfileLogic {
	return &GetUniqueProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUniqueProfileLogic) GetUniqueProfile() (*types.GetUniqueProfileRes, error) {
	resp := &types.GetUniqueProfileRes{
		Id:       l.ctx.Value("id").(string),
		UserId:   l.ctx.Value("userId").(string),
		Name:     l.ctx.Value("name").(string),
		ImageUrl: l.ctx.Value("imageUrl").(string),
		Email:    l.ctx.Value("email").(string),
	}
	return resp, nil
}

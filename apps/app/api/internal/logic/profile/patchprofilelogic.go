package profile

import (
	"context"
	"errors"
	profileRPC "github.com/POABOB/go-discord/apps/profile/rpc/pb/rpc"

	"github.com/POABOB/go-discord/apps/app/api/internal/svc"
	"github.com/POABOB/go-discord/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PatchProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPatchProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PatchProfileLogic {
	return &PatchProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PatchProfileLogic) PatchProfile(req *types.PatchProfileReq) error {
	payloadId := l.ctx.Value("id").(string)
	if req.Id != payloadId {
		return errors.New("cannot patch other profile")
	}

	// Call RPC
	if _, err := l.svcCtx.Profile.PatchProfile(l.ctx, &profileRPC.PatchProfileReq{
		Id:       payloadId,
		Name:     req.Name,
		Email:    req.Email,
		ImageUrl: req.ImageUrl,
	}); err != nil {
		return err
	}

	return nil
}

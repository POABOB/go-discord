package profile

import (
	"context"
	profileRPC "github.com/POABOB/go-discord/apps/profile/rpc/pb/rpc"
	"google.golang.org/grpc/status"

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
		return status.Error(403, "cannot patch other profile")
	}

	// Call RPC
	if _, err := l.svcCtx.Profile.PatchProfile(l.ctx, &profileRPC.PatchProfileReq{
		Id:       payloadId,
		Name:     req.Name,
		Email:    req.Email,
		ImageUrl: req.ImageUrl,
	}); err != nil {
		return status.Error(400, err.Error())
	}

	return nil
}

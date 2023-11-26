package profile

import (
	"context"
	"errors"
	"github.com/POABOB/go-discord/apps/app/api/internal/svc"
	"github.com/POABOB/go-discord/apps/app/api/internal/types"
	profileRPC "github.com/POABOB/go-discord/apps/profile/rpc/pb/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProfileLogic {
	return &DeleteProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteProfileLogic) DeleteProfile(req *types.DeleteProfileReq) error {
	payloadId := l.ctx.Value("id").(string)
	if req.Id != payloadId {
		return errors.New("cannot delete other profile")
	}

	// Call RPC
	if _, err := l.svcCtx.Profile.DeleteProfile(l.ctx, &profileRPC.DeleteProfileReq{Id: payloadId}); err != nil {
		return err
	}

	return nil
}

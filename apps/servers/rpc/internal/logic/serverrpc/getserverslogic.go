package serverrpclogic

import (
	"context"
	"github.com/POABOB/go-discord/prisma/db"

	"github.com/POABOB/go-discord/apps/servers/rpc/internal/svc"
	"github.com/POABOB/go-discord/apps/servers/rpc/pb/rpc"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetServersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetServersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetServersLogic {
	return &GetServersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetServersLogic) GetServers(in *rpc.GetServersReq) (*rpc.GetServersRes, error) {
	ctx := context.Background()

	var input rpc.GetServersReq
	if *in.Id != "" {
		input.Id = in.Id
	}
	if *in.ProfileId != "" {
		input.ProfileId = in.ProfileId
	}
	if *in.Name != "" {
		input.Name = in.Name
	}
	if *in.InviteCode != "" {
		input.InviteCode = in.InviteCode
	}

	// get
	serverModel, err := l.svcCtx.PrismaClient.Server.
		FindMany(
			db.Server.ID.EqualsIfPresent(input.Id),
			db.Server.ProfileID.EqualsIfPresent(input.ProfileId),
			db.Server.Name.EqualsIfPresent(input.Name),
			db.Server.InviteCode.EqualsIfPresent(input.InviteCode),
		).Take(int(in.PagePerNum)).Skip(int(in.PagePerNum * (in.Page - 1))).Exec(ctx)
	if err != nil {
		return &rpc.GetServersRes{}, err
	}

	var servers rpc.GetServersRes
	if err := copier.Copy(&servers.Servers, &serverModel); err != nil {
		return &rpc.GetServersRes{}, err
	}
	return &servers, nil
}

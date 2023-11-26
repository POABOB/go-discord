package serverrpclogic

import (
	"context"
	"github.com/POABOB/go-discord/prisma/db"
	"github.com/google/uuid"

	"github.com/POABOB/go-discord/apps/servers/rpc/internal/svc"
	"github.com/POABOB/go-discord/apps/servers/rpc/pb/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostServerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostServerLogic {
	return &PostServerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PostServerLogic) PostServer(in *rpc.PostServerReq) (*rpc.Empty, error) {
	ctx := context.Background()

	// create
	uid, err := uuid.NewUUID()
	serverUuid, err := uuid.NewUUID()
	if err != nil {
		return &rpc.Empty{}, err
	}
	postServer := l.svcCtx.PrismaClient.Server.CreateOne(
		db.Server.Name.Set(in.Name),
		db.Server.ImageURL.Set(in.ImageUrl),
		db.Server.InviteCode.Set(uid.String()),
		db.Server.Profile.Link(
			db.Profile.ID.Equals(in.ProfileId),
		),
		db.Server.ID.Set(serverUuid.String()),
	).Tx()

	postMember := l.svcCtx.PrismaClient.Member.CreateOne(
		db.Member.Profile.Link(db.Profile.ID.Equals(in.ProfileId)),
		db.Member.Server.Link(db.Server.ID.Equals(serverUuid.String())),
		db.Member.Role.Set(db.MemberRoleADMIN),
	).Tx()
	if err := l.svcCtx.PrismaClient.Prisma.Transaction(postServer, postMember).Exec(ctx); err != nil {
		return &rpc.Empty{}, err
	}

	return &rpc.Empty{}, nil
}

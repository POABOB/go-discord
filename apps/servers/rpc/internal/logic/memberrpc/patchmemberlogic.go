package memberrpclogic

import (
	"context"
	"github.com/POABOB/go-discord/apps/servers/rpc/internal/logic/serverrpc"
	"github.com/POABOB/go-discord/apps/servers/rpc/internal/svc"
	"github.com/POABOB/go-discord/apps/servers/rpc/pb/rpc"
	"github.com/POABOB/go-discord/prisma/db"

	"github.com/zeromicro/go-zero/core/logx"
)

type PatchMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPatchMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PatchMemberLogic {
	return &PatchMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PatchMemberLogic) PatchMember(in *rpc.PatchMemberReq) (*rpc.GetServerRes, error) {
	ctx := context.Background()

	memberType := map[string]db.MemberRole{
		"ADMIN":     db.MemberRoleADMIN,
		"MODERATOR": db.MemberRoleMODERATOR,
		"GUEST":     db.MemberRoleGUEST,
	}

	patchMember := l.svcCtx.PrismaClient.Member.
		FindUnique(db.Member.ID.Equals(in.Id)).
		Update(db.Member.Role.Set(memberType[in.Role])).Tx()
	if err := l.svcCtx.PrismaClient.Prisma.Transaction(patchMember).Exec(ctx); err != nil {
		return nil, err
	}

	getUniqueServer := serverrpclogic.NewGetUniqueServerLogic(l.ctx, l.svcCtx)
	return getUniqueServer.GetUniqueServer(&rpc.GetServerReq{Id: in.ServerId, ProfileId: in.ProfileId})
}

package memberrpclogic

import (
	"context"
	"errors"
	"github.com/POABOB/go-discord/prisma/db"
	"github.com/jinzhu/copier"

	"github.com/POABOB/go-discord/apps/servers/rpc/internal/svc"
	"github.com/POABOB/go-discord/apps/servers/rpc/pb/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type InviteMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInviteMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InviteMemberLogic {
	return &InviteMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InviteMemberLogic) InviteMember(in *rpc.InviteMemberReq) (*rpc.GetServerRes, error) {
	ctx := context.Background()
	// 用 inviteCode 抓 Server
	serverModel, err := l.svcCtx.PrismaClient.Server.
		FindUnique(db.Server.InviteCode.Equals(in.InviteCode)).
		With(
			db.Server.Members.Fetch().
				With(db.Member.Profile.Fetch()),
		).
		Exec(ctx)
	if errors.Is(err, db.ErrNotFound) {
		return nil, errors.New("無效的伺服器邀請碼！")
	} else if err != nil {
		return nil, err
	}

	serverId := serverModel.InnerServer.ID
	members := serverModel.RelationsServer.Members
	var member db.MemberModel
	for _, v := range members {
		if v.ProfileID == in.ProfileId {
			member = v
		}
	}

	var server rpc.GetServerRes
	server.Server = &rpc.Server{}
	if err := copier.Copy(&server.Server, serverModel.InnerServer); err != nil {
		return &rpc.GetServerRes{}, err
	}
	if member.ID != "" {
		return &server, nil
	}

	postMember := l.svcCtx.PrismaClient.Member.CreateOne(
		db.Member.Profile.Link(db.Profile.ID.Equals(in.ProfileId)),
		db.Member.Server.Link(db.Server.ID.Equals(serverId)),
		db.Member.Role.Set(db.MemberRoleGUEST),
	).Tx()
	if err := l.svcCtx.PrismaClient.Prisma.Transaction(postMember).Exec(ctx); err != nil {
		return nil, err
	}

	return &server, nil
}

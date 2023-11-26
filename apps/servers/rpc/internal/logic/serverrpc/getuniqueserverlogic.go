package serverrpclogic

import (
	"context"
	"github.com/POABOB/go-discord/apps/servers/rpc/internal/svc"
	"github.com/POABOB/go-discord/apps/servers/rpc/pb/rpc"
	"github.com/POABOB/go-discord/prisma/db"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUniqueServerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUniqueServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUniqueServerLogic {
	return &GetUniqueServerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUniqueServerLogic) GetUniqueServer(in *rpc.GetServerReq) (*rpc.GetServerRes, error) {
	ctx := context.Background()

	// get unique
	serverModel, err := l.svcCtx.PrismaClient.Server.
		FindFirst(
			db.Server.ID.Equals(in.Id),
		).
		With(
			db.Server.Channels.Fetch().With(),
			db.Server.Members.Fetch().OrderBy(db.Member.CreatedAt.Order(db.SortOrderAsc)).
				With(
					db.Member.Profile.Fetch(),
				),
		).Exec(ctx)
	if err != nil {
		return &rpc.GetServerRes{}, err
	}
	var server rpc.GetServerRes
	server.Server = &rpc.Server{}
	server.Relation = &rpc.Relation{}
	if err := copier.Copy(&server.Server, serverModel.InnerServer); err != nil {
		return &rpc.GetServerRes{}, err
	}
	memberType := map[db.MemberRole]string{
		db.MemberRoleADMIN:     "ADMIN",
		db.MemberRoleMODERATOR: "MODERATOR",
		db.MemberRoleGUEST:     "GUEST",
	}
	for _, v := range serverModel.RelationsServer.Members {
		temp := rpc.Member{
			Id:        v.InnerMember.ID,
			Role:      memberType[v.InnerMember.Role],
			ProfileId: v.RelationsMember.Profile.InnerProfile.ID,
			Name:      v.RelationsMember.Profile.InnerProfile.Name,
			ImageUrl:  v.RelationsMember.Profile.InnerProfile.ImageURL,
			Email:     v.RelationsMember.Profile.InnerProfile.Email,
		}
		server.Relation.Members = append(server.Relation.Members, &temp)
	}
	channelType := map[db.ChannelType]string{
		db.ChannelTypeTEXT:  "TEXT",
		db.ChannelTypeAUDIO: "AUDIO",
		db.ChannelTypeVIDEO: "VIDEO",
	}
	for _, v := range serverModel.RelationsServer.Channels {
		temp := rpc.Channel{
			Id:       v.InnerChannel.ID,
			Name:     v.InnerChannel.Name,
			Type:     channelType[v.InnerChannel.Type],
			ServerId: v.InnerChannel.ServerID,
		}
		server.Relation.Channels = append(server.Relation.Channels, &temp)
	}

	//fmt.Println(server.Relation.Members)
	return &server, nil
}

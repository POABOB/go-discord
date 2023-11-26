package auth

import (
	"context"
	profileRPC "github.com/POABOB/go-discord/apps/profile/rpc/pb/rpc"
	"github.com/golang-jwt/jwt/v4"
	"time"

	"github.com/POABOB/go-discord/apps/app/api/internal/svc"
	"github.com/POABOB/go-discord/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (*types.LoginRes, error) {
	claims, err := l.svcCtx.Clerk.DecodeToken(req.Token)
	if err != nil {
		return nil, err
	}

	// Call RPC
	myClaim := claims.Extra
	temp, err := l.svcCtx.Profile.GetUniqueProfileOrCreate(l.ctx, &profileRPC.GetUniqueProfileOrCreateReq{
		UserId:   myClaim["userId"].(string),
		Name:     myClaim["name"].(string),
		ImageUrl: myClaim["imageUrl"].(string),
		Email:    myClaim["email"].(string),
	})
	if err != nil {
		return nil, err
	}

	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, time.Now().Unix(), l.svcCtx.Config.Auth.AccessExpire, temp.Id, myClaim)
	if err != nil {
		return nil, err
	}

	return &types.LoginRes{Id: temp.Id, Token: jwtToken}, nil
}

func (l *LoginLogic) getJwtToken(secret string, nowDate int64, accessExpire int64, Id string, myClaim map[string]interface{}) (jwtToken string, err error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = nowDate + accessExpire
	claims["iat"] = nowDate
	claims["id"] = Id
	for k, v := range myClaim {
		claims[k] = v
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	jwtToken, err = token.SignedString([]byte(secret))
	return
}

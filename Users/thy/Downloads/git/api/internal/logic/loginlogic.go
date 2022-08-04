package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"goZeroDemo/Users/thy/Downloads/git/api/internal/svc"
	"goZeroDemo/Users/thy/Downloads/git/api/internal/types"
	"strings"
	"time"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type User struct {
	Id         int64
	Name       string // 用户姓名
	Gender     int64  // 用户性别
	Mobile     string // 用户电话
	Password   string // 用户密码
	CreateTime time.Time
	UpdateTime time.Time
}

func (l *LoginLogic) Login(req *types.LoginReq) (*types.LoginReply, error) {
	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.New("参数错误")
	}

	var userInfo User
	//switch err {
	//case nil:
	//case sqlc.ErrNotFound:
	//	return nil, errors.New("用户名不存在")
	//default:
	//	return nil, err
	//}
	//
	//if userInfo.Password != req.Password {
	//	return nil, errors.New("用户密码不正确")
	//}

	// ---start---
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, userInfo.Id)
	if err != nil {
		return nil, err
	}
	// ---end---

	return &types.LoginReply{
		Id:   userInfo.Id,
		Name: userInfo.Name,
		//Gender:       string(userInfo.Gender),
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

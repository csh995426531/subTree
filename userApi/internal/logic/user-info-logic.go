package logic

import (
	"context"
	"subModule/userRpc/userrpcclient"

	"subModule/userApi/internal/svc"
	"subModule/userApi/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserInfoLogic {
	return UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req types.InfoReq) (*types.InfoReply, error) {
	// todo: add your logic here and delete this line

	resp, err := l.svcCtx.UserRpc.Info(l.ctx, &userrpcclient.InfoRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.InfoReply{
		Id: resp.Id,
		Name: resp.Name,
	}, nil
}

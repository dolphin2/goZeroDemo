package svc

import (
	"goZeroDemo/Users/thy/Downloads/git/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
	//	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	//sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		//UserModel: model.NewUserModel(nil, c.CacheRedis),
	}
}

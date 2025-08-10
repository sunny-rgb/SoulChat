package user

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-realworld/internal/biz/user"
	"kratos-realworld/internal/data_struct"
)

type UserLogRepo struct {
	data *data_struct.Data
	log  *log.Helper
}

func NewUserLogRepo(data *data_struct.Data, logger log.Logger) *UserLogRepo {
	return &UserLogRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (r *UserLogRepo) CreateUser(ctx context.Context, u *user.UserLog) error {
	db := r.data.Db
	if err := db.Model(u).Create(u).Error; err != nil {
		r.log.Errorf("CreateUser fail: %v", err)
		return fmt.Errorf("CreateUser fail: %v", err)
	}
	r.log.Infof("insert success")
	return nil
}

func (r *UserLogRepo) UpdateByCache(u *user.UserLog) error {
	// TODO: Implement cache update logic
	r.log.Infof("UpdateByCache called for user: %s", u.UserName)
	return nil
}

func (r *UserLogRepo) Load(u *user.UserLog) error {
	// TODO: Implement load logic
	r.log.Infof("Load called for user: %s", u.UserName)
	return nil
}

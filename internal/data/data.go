package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"kratos-realworld/internal/data/user"
	"kratos-realworld/internal/data_struct"
)

var ProviderSet = wire.NewSet(
	user.NewUserLogRepo,
	data_struct.NewData,
	data_struct.NewCache,
	data_struct.NewDatabase,
	data_struct.NewTransaction,
)

// NewUserLogRepoProvider creates a UserLogRepo with dependencies
func NewUserLogRepoProvider(data *data_struct.Data, logger log.Logger) *user.UserLogRepo {
	return user.NewUserLogRepo(data, logger)
}

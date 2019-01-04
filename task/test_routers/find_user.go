package test_routers

import (
	"context"
	"github.com/johnnyeven/libtools/courier/httpx"
	"github.com/johnnyeven/libtools/courier"
)

func init() {
	Router.Register(courier.NewRouter(FindUser{}))
}

// 查找用户
type FindUser struct {
	httpx.MethodGet
	// 用户ID
	UserID uint64 `name:"userId,string" in:"path"`
}

func (req FindUser) Path() string {
	return "/:userId"
}

func (req FindUser) Output(ctx context.Context) (result interface{}, err error) {
	return struct {
		UserID   uint64
		UserName string
	}{
		req.UserID,
		"johnnyeven",
	}, nil
}

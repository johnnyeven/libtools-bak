package client

import (
	"github.com/profzone/libtools/courier"
)

type IRequest interface {
	Do() courier.Result
}

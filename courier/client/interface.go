package client

import (
	"profzone/libtools/courier"
)

type IRequest interface {
	Do() courier.Result
}

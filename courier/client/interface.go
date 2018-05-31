package client

import (
	"golib/tools/courier"
)

type IRequest interface {
	Do() courier.Result
}

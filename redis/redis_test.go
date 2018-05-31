package redis_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"golib/tools/redis"
)

func TestRedis_GetSet(t *testing.T) {
	tt := assert.New(t)

	r := redis.Redis{
		Host: "staging.g7pay.chinawayltd.com",
		Port: 36379,
	}
	r.MarshalDefaults(&r)

	r.Init()

	err := r.GetCache().Set("test", []byte("test"), 7200*time.Second)
	tt.Nil(err)

	bytes, err := r.GetCache().Get("test")
	tt.Nil(err)

	tt.Equal("test", string(bytes))
}

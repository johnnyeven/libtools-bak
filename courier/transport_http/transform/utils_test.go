package transform

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTagValueAndFlags(t *testing.T) {
	tt := assert.New(t)

	cases := []struct {
		tag   string
		v     string
		flags TagFlags
	}{
		{
			tag:   "json",
			v:     "json",
			flags: TagFlags{},
		},
		{
			tag: "json,omitempty",
			v:   "json",
			flags: TagFlags{
				"omitempty": true,
			},
		},
	}

	for _, tc := range cases {
		v, flags := GetTagValueAndFlags(tc.tag)
		tt.Equal(tc.v, v)
		tt.Equal(tc.flags, flags)
	}
}

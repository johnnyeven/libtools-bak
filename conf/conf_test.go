package conf_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"golib/tools/conf"
)

type E struct {
	E int
}

func (e E) MarshalDefaults(v interface{}) {
	if target, ok := v.(*E); ok {
		if target.E == 0 {
			target.E = 4
		}
	}
}

type D struct {
	E
	F float64
	G string
}

type StructStruct struct {
	A  int
	B  bool
	C  uint
	D  D
	DP *D
	H  []string
}

func TestUnmarshal(t *testing.T) {
	config := &StructStruct{
		A: 1,
		B: false,
		C: 2,
		D: D{
			E: E{},
		},
		DP: &D{
			E: E{
				E: 4,
			},
			F: 2.0,
			G: "abc",
		},
		H: []string{"def"},
	}

	os.Setenv("TEST_A", "2")
	os.Setenv("TEST_B", "true")
	os.Setenv("TEST_D_E", "5")
	os.Setenv("TEST_DP_E", "5")
	os.Setenv("TEST_H", "gds,123")

	conf.Unmarshal(reflect.ValueOf(config), "TEST")

	assert.Equal(t, &StructStruct{
		A: 2,
		B: true,
		C: 2,
		D: D{
			E: E{
				E: 5,
			},
		},
		DP: &D{
			E: E{
				E: 5,
			},
			F: 2.0,
			G: "abc",
		},
		H: []string{"gds", "123"},
	}, config)
}

package gen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseEnum(t *testing.T) {
	doc, hasEnum := ParseEnum("swagger:enum \nasdasdasdad")
	assert.Equal(t, "asdasdasdad", doc)
	assert.Equal(t, true, hasEnum)
}

func TestParseStrfmt(t *testing.T) {
	doc, fmtName := ParseStrfmt("swagger:strfmt date-time\nasdasdasdad")
	assert.Equal(t, "asdasdasdad", doc)
	assert.Equal(t, "date-time", fmtName)
}

func TestGetCommonValidations(t *testing.T) {
	c := getCommonValidations("@int[1,2]")
	assert.NotNil(t, c.Minimum)
	assert.NotNil(t, c.Maximum)
	assert.Equal(t, false, c.ExclusiveMinimum)
	assert.Equal(t, false, c.ExclusiveMaximum)
}

func TestGetCommonValidationsWithExclusive(t *testing.T) {
	c := getCommonValidations("@int(1,2)")
	assert.NotNil(t, c.Minimum)
	assert.NotNil(t, c.Maximum)
	assert.Equal(t, true, c.ExclusiveMinimum)
	assert.Equal(t, true, c.ExclusiveMaximum)
}

func TestGetCommonValidationsWithEnum(t *testing.T) {
	c := getCommonValidations("@int{1,2}")
	assert.Equal(t, c.Enum, []interface{}{int64(1), int64(2)})
}

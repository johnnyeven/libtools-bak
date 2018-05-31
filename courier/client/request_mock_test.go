package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnMarshalRequestID(t *testing.T) {
	tt := assert.New(t)
	mock := Mock("a").For("b.some", MockData{
		Data: []byte(`{"a":1}`),
	})
	parsedMock, err := ParseMockID("a", mock.RequestID())
	tt.Nil(err)
	tt.Equal(mock, parsedMock)
}

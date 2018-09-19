package service_test

import (
	"testing"

	"github.com/profzone/libtools/service"
)

func TestNew(t *testing.T) {
	serve := service.New("test")
	if serve.Name != "test" {
		t.Fatalf("name is not test")
	}
}

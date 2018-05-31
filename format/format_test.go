package format_test

import (
	"fmt"
	"testing"

	"golib/tools/format"
)

func TestProcess(t *testing.T) {
	result, _ := format.Process("format2_test.go", []byte(`
	package format

	import (
		testing "testing"
		"fmt"

		"golib/tools/format"
		"golib/tools/gin_app"

		"github.com/davecgh/go-spew/spew"
	)

	func Test(t *testing.T) {
		spew.Dump(gin_app.REQUEST_ID_NAME, format.Test)
	}
	`))

	fmt.Println(string(result))
}

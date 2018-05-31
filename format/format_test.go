package format_test

import (
	"fmt"
	"testing"

	"profzone/libtools/format"
)

func TestProcess(t *testing.T) {
	result, _ := format.Process("format2_test.go", []byte(`
	package format

	import (
		testing "testing"
		"fmt"

		"profzone/libtools/format"
		"profzone/libtools/gin_app"

		"github.com/davecgh/go-spew/spew"
	)

	func Test(t *testing.T) {
		spew.Dump(gin_app.REQUEST_ID_NAME, format.Test)
	}
	`))

	fmt.Println(string(result))
}

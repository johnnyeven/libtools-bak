// package
package fixtures

import (
	// fmt
	"fmt"
	"time"
)

// type
type Date time.Time

// type
type Test struct {
	// TestFieldString
	String string
	// TestFieldInt
	Int int
	// TestFieldBool
	Bool bool
	// TestFieldDate
	Date Date
}

// type
type (
	Test2 struct {
		// TestFieldString
		String string
		// TestFieldInt
		Int int
		// TestFieldBool
		Bool bool
	}
)

// func
func (t Test2) Recv() {

}

// var
var test = Test{
	String: "",
	Int:    1 + 1,
	Bool:   true,
}

// var
var (
	// test2
	test2 = Test{
		String: "",
		Int:    1,
		Bool:   true,
	}
	// test2
	test3 = Test{
		String: "",
		Int:    1,
		Bool:   true,
	}
)

// func
func Print(a string, b string) string {
	return a + b
}

// func
func fn() {
	// Call
	res := Print("", "")
	if res != "" {
		// print
		fmt.Println(res)
	}
}

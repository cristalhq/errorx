package errorx_test

import (
	"fmt"
	"os"

	"github.com/cristalhq/errorx"
)

func ExampleNew() {
	err := errorx.New("this is the error")
	err2 := errorx.New("this is the error")

	if err == err2 {
		panic("should be different")
	}

	fmt.Println(err)
	fmt.Println(err2)

	// Output:
	// this is the error
	// this is the error
}

func ExampleNewf() {
	err := errorx.Newf("this is the error, code: %d", 123)
	err2 := errorx.Newf("this is the error, code: %d", 123)

	if err == err2 {
		panic("should be different")
	}

	fmt.Println(err)
	fmt.Println(err2)

	// Output:
	// this is the error, code: 123
	// this is the error, code: 123
}

func ExampleNewf_withW() {
	err := errorx.New("this is the error")
	errFull := errorx.Newf("caused by: %w", err)

	if err == errFull {
		panic("should be different")
	}

	fmt.Println(err)
	fmt.Println(errFull)

	// Output:
	// this is the error
	// caused by: this is the error
}

func ExampleIsAny() {
	err := os.ErrPermission

	if errorx.IsAny(err, os.ErrNotExist, os.ErrPermission) {
		fmt.Println("it's not DNS")
	}

	// Output:
	// it's not DNS
}

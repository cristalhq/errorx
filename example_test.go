package errorx_test

import (
	"fmt"
	"os"
	"time"

	"github.com/cristalhq/errorx"
)

func Example() {
	func() error {
		err := errorx.Newf("this is the error")
		if err != nil {
			return errorx.Wrapf(err, "something happened")
		}

		errAt := errorx.Newf("happened at: %s", time.Now())
		if errAt != nil {
			return errorx.Trace(err)
		}

		if errorx.Tracing() {
			println("error tracing is enabled")
		}
		return nil
	}()
}

func ExampleNewf() {
	err := errorx.Newf("this is the error")
	err2 := errorx.Newf("this is the error, with code: %d", 123)
	err3 := errorx.Newf("this is the error, with code: %d", 123)
	errFull := errorx.Newf("caused by: %w", err3)

	if err2 == err3 {
		panic("should be different")
	}

	fmt.Println(err)
	fmt.Println(err2)
	fmt.Println(err3)
	fmt.Println(errFull)

	// Output:
	// this is the error
	// this is the error, with code: 123
	// this is the error, with code: 123
	// caused by: this is the error, with code: 123
}

func ExampleIsAny() {
	err := os.ErrPermission

	if errorx.IsAny(err, os.ErrNotExist, os.ErrPermission) {
		fmt.Println("it's not DNS")
	}

	// Output:
	// it's not DNS
}

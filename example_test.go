package errorx_test

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/cristalhq/errorx"
)

func Example() {
	err := func() error {
		err := errorx.Newf("this is the error")
		if err != nil {
			return errorx.Wrapf(err, "something happened")
		}

		errAt := errorx.Newf("happened at: %s", time.Now())
		if errAt != nil {
			return errorx.Trace(errAt)
		}

		if errorx.Tracing() {
			println("error tracing is enabled")
		}
		return nil
	}()

	fmt.Println(err.Error())

	// Output:
	// something happened: this is the error
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

func ExampleTrace() {
	errorx.EnableTrace()
	defer errorx.DisableTrace()

	err := errors.New("root error")
	err2 := errorx.Trace(err)

	_ = err2

	// Output:
}

func ExampleWrapf() {
	// Output:
}

func ExampleIs() {
	err := os.ErrPermission

	if errorx.Is(err, os.ErrNotExist) {
		panic("not this case")
	}

	if errorx.Is(err, os.ErrNotExist, os.ErrPermission) {
		fmt.Println("it's not DNS")
	}

	if errorx.Is(err, os.ErrPermission) {
		fmt.Println("oh no it's permissions")
	}

	// Output:
	// it's not DNS
	// oh no it's permissions
}

func ExampleAs() {
}

func ExampleInto() {
}

func ExampleUnwrap() {
}

func Example_multiErrors() {
}

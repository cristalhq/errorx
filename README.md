# errorx

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![version-img]][version-url]

TODO

## Rationale

TODO

## Features

* Simple and easy.
* Safe & fast.
* Tested.
* Dependency-free.

See [these docs][pkg-url] or [GUIDE.md](GUIDE.md) for more details.

## Install

Go version 1.18+

```
go get github.com/cristalhq/errorx
```

## Example

```go
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
```

See examples: [example_test.go](example_test.go).

## License

[MIT License](LICENSE).

[build-img]: https://github.com/cristalhq/errorx/workflows/build/badge.svg
[build-url]: https://github.com/cristalhq/errorx/actions
[pkg-img]: https://pkg.go.dev/badge/cristalhq/errorx
[pkg-url]: https://pkg.go.dev/github.com/cristalhq/errorx
[version-img]: https://img.shields.io/github/v/release/cristalhq/errorx
[version-url]: https://github.com/cristalhq/errorx/releases

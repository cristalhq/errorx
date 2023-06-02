package errorx

type Formatter interface {
	error
	FormatError(p Printer) (next error)
}

type Printer interface {
	Print(args ...any)
	Printf(format string, args ...any)
	Detail() bool
}

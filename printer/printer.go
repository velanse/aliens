package printer

import "fmt"

type Printer interface {
	Printf(format string, a ...interface{})
	Println(a ...interface{})
	Debug(format string, a ...interface{})
}

type StdoutPrinter struct {
	debug bool
}

func NewStdoutPrinter(debug bool) Printer {
	return &StdoutPrinter{debug: debug}
}

func (p StdoutPrinter) Printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func (p StdoutPrinter) Println(a ...interface{}) {
	fmt.Println(a...)
}

func (p StdoutPrinter) Debug(format string, a ...interface{}) {
	if p.debug {
		fmt.Printf(format, a...)
	}
}

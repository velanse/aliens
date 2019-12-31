package printer

import "fmt"

type MockPrinter struct {
	Messages []string
}

func (p *MockPrinter) Printf(format string, a ...interface{}) {
	p.Messages = append(p.Messages, fmt.Sprintf(format, a...))
}

func (p *MockPrinter) Println(a ...interface{}) {
	p.Messages = append(p.Messages, fmt.Sprintln(a...))
}

func (p *MockPrinter) Debug(format string, a ...interface{}) {

}

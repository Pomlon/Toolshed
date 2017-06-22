package pomlog

import "fmt"

type LogConsole struct {
}

func (l LogConsole) Setup() {}

func (l LogConsole) Teardown() {}

func (l LogConsole) LogOutput(msg []string) {
	for _, ting := range msg {
		fmt.Print(ting + " ")
	}
	fmt.Println()
}

package pomlog

import "fmt"

func Spawn() Logger {
	return Logger{
		level: 1,
		logOuts: []LogOutputter{
			LogConsole{},
		},
	}
}

var Error = 0
var Info = 1
var Debug = 2

type LogOutputter interface {
	Setup()
	Teardown()
	LogOutput([]string)
}

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

type LogFormatter interface {
	Format(msg []string) string
}

type Logger struct {
	level   int
	logOuts []LogOutputter
}

//SetLevel sets the logging level. 0 - error, 1 - info, 2 - debug
func (l Logger) SetLevel(logLevel int) {
	if logLevel > 2 {
		logLevel = 2
	} else if logLevel < 0 {
		logLevel = 0
	}
	l.level = logLevel
}

func (l Logger) Error(msg ...string) {
	l.log(Error, msg)
}

func (l Logger) Info(msg ...string) {
	l.log(Info, msg)
}

func (l Logger) Debug(msg ...string) {
	l.log(Debug, msg)
}

func (l Logger) log(lvl int, msg []string) {
	if lvl <= l.level {
		for _, logOut := range l.logOuts {

			logOut.LogOutput(msg)

		}
	}
}

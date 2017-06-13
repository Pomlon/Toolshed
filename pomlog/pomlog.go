package pomlog

import "fmt"

func Spawn() Logger {
	return Logger{
		level: 0,
		logOuts: []LogOutputter{
			LogConsole{},
		},
	}
}

var Info = 0
var Warn = 1
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
	Format(msg []string)
}

type Logger struct {
	level   int
	logOuts []LogOutputter
}

//SetLevel sets the logging level. 0 - info, 1 - warn, 2 - debug
func (l Logger) SetLevel(logLevel int) {
	if logLevel > 2 {
		logLevel = 2
	} else if logLevel < 0 {
		logLevel = 0
	}
	l.level = logLevel
}

func (l Logger) Info(msg ...string) {
	l.log(0, msg)
}

func (l Logger) Warn(msg ...string) {
	l.log(1, msg)
}

func (l Logger) Debug(msg ...string) {
	l.log(2, msg)
}

func (l Logger) log(lvl int, msg []string) {
	if lvl <= l.level {
		for _, logOut := range l.logOuts {
			//logOut.Setup()
			logOut.LogOutput(msg)
			//logOut.Teardown()
		}
	}
}

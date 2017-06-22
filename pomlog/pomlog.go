package pomlog

func Spawn() Logger {
	return Logger{
		level: 1,
		LogOuts: []LogOutputter{
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

type LogFormatter interface {
	Format(msg []string) []string
}

type Logger struct {
	level     int
	LogOuts   []LogOutputter
	LogFormat LogFormatter
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
		for _, logOut := range l.LogOuts {
			if l.LogFormat != nil {
				logOut.LogOutput(l.LogFormat.Format(msg))
			} else {

				logOut.LogOutput(msg)
			}

		}
	}
}

package pomlog

import (
	"io/ioutil"

	"github.com/Pomlon/Toolshed"
)

type LogToFile struct {
	Filename string
}

func (l LogToFile) Setup() {}

func (l LogToFile) Teardown() {}

func (l LogToFile) LogOutput(msg []string) {
	str := ""
	for _, ting := range msg {
		str += ting + " "
	}
	str += "\n"
	bytes := []byte(str)

	err := ioutil.WriteFile(l.Filename, bytes, 0644)
	Toolshed.ErrLog(err)
}

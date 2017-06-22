package pomlog

import (
	"time"
)

var DateExtended = "Mon 02-01-2006 15:04:05"
var DateFull = "02-01-2006 15:04:05"
var DateFullShortened = "02-01-06 15:04:05"

type StdFormat struct {
	LogFormatter
	Date string
	Tags []string
}

func (f StdFormat) Format(msg []string) []string {
	out := ""

	if f.Date != "" {
		out += time.Now().Format(f.Date) + " "
	}

	if len(f.Tags) > 0 {
		out += " ["
		one := true
		for _, tag := range f.Tags {
			if one == true {
				out += tag
				one = false
			} else {
				out += "/" + tag
			}
		}
		out += "] "
	}

	for _, txt := range msg {
		out += txt + " "
	}

	return []string{out}
}

func (f StdFormat) AddTags(s ...string) {
	for _, txt := range s {
		f.Tags = append(f.Tags, txt)
	}
}

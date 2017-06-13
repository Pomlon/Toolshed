//Package Toolshed contains various misc tools which extend the functionality of GO, and often bring in familiar utils from Python
package Toolshed

import (
	"fmt"
	"strings"
)

//JoinDelimiter specifies the delimiter by which to join strings back after Frmtting them.
//You can either change this, or create a new Stronk and set its' delim field
var JoinDelimiter = " "

type Stronk struct {
	s     string
	delim string
}

//Frmt mimics old Pythons string.format functions behavior. This is NOT meant to use by itself!
//This thing only creates a struct containing a string. The goodies are strapped to that struct itself.
func Frmt(str string) Stronk {
	return Stronk{s: str, delim: JoinDelimiter}
}

func (str Stronk) F(args ...string) string {
	foo := strings.SplitAfter(str.s, "{}")
	replaceCounter := 0

	if foo[len(foo)-1] == "" {
		foo = foo[:len(foo)-1]
	}
	//BUG(Pomlon) This means F will replace one more occurence of {} (if present) with the last arg.
	if len(foo) > len(args)+1 || len(foo) < len(args) {
		fmt.Printf("Replace param count doesn't match replace wildcards %d to %d \n string: %s \n", len(foo), len(args), str.s)
		return ""
	}

	for i, spl := range foo {
		tmp := strings.TrimSpace(spl)
		fmt.Println("replace counter ", replaceCounter)
		foo[i] = strings.Replace(tmp, "{}", args[replaceCounter], 1)
		if len(args)-1 > replaceCounter {
			replaceCounter++
		}
	}

	return strings.Join(foo, str.delim)
}

//Abs takes int, returns absolute.
func Abs(i int) int {
	if i < 0 {
		return -i
	}
	if i == 0 {
		return 0
	}
	return i
}

//ErrPanic freaks out on non-empty error
func ErrPanic(e error) {
	if e != nil {
		panic(e)
	}
}

//ErrLog simply logs an error to stdout
func ErrLog(e error) {
	if e != nil {
		fmt.Println("Error!", e)
	}
}

package main

import "fmt"

//LogLevel is a logging level
type LogLevel uint8

//possible log levels
const (
	DebugLevel LogLevel = iota + 1
	WarningLevel
	ErrorLevel
)

//string implements the fmt.Stringer interface
func (l LogLevel) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case WarningLevel:
		return "warning"
	case ErrorLevel:
		return "error"
	}
	return fmt.Sprintf("unknown log level: %d", l)
}

func main() {
	fmt.Println(WarningLevel)

	lvl := LogLevel(19)
	fmt.Println(lvl)
}

/*
OUTPUT
warning
unknown log level: 19
*/

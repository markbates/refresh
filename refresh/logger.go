package refresh

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

const lformat = "=== %s ==="

type Logger struct {
	log *log.Logger
}

func NewLogger(c *Configuration) *Logger {
	color.NoColor = !c.EnableColors
	return &Logger{
		log: log.New(os.Stdout, "refresh: ", log.LstdFlags),
	}
}

func (l *Logger) Success(msg interface{}, args ...interface{}) {
	l.log.Print(color.GreenString(fmt.Sprintf(lformat, msg), args...))
}

func (l *Logger) Error(msg interface{}, args ...interface{}) {
	l.log.Print(color.RedString(fmt.Sprintf(lformat, msg), args...))
}

func (l *Logger) Print(msg interface{}, args ...interface{}) {
	l.log.Printf(fmt.Sprintf(lformat, msg), args...)
}

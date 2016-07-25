package refresh

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

const lformat = "=== %s ===\n"

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
	color.Green(fmt.Sprintf(lformat, msg), args...)
}

func (l *Logger) Error(msg interface{}, args ...interface{}) {
	color.Red(fmt.Sprintf(lformat, msg), args...)
}

func (l *Logger) Print(msg interface{}, args ...interface{}) {
	l.log.Printf(fmt.Sprintf(lformat, msg), args...)
}

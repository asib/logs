package logs

import (
	"time"
)

type LogLevel int

const (
	Info LogLevel = iota
	Warning
	Error
)

type LogMsg struct {
	Level LogLevel
	Msg   string
}

type LogStack struct {
	Logger *Logger
	Stack  []*LogMsg
}

func (ls *LogStack) Add(lvl LogLevel, msg string) {
	ls.Stack = append(ls.Stack,
		&LogMsg{
			Level: lvl,
			Msg:   msg,
		})
}

func (ls *LogStack) PrintStack() {
	ls.Logger.mu.Lock()
	if time.Since(ls.Logger.Last) > ls.Logger.Timeout {
		for _, log := range ls.Stack {
			switch log.Level {
			case Info:
				ls.Logger.Info.Println(log.Msg)
			case Warning:
				ls.Logger.Warning.Println(log.Msg)
			case Error:
				ls.Logger.Error.Println(log.Msg)
			default:
			}
		}
		ls.Logger.Last = time.Now()
	}
	ls.Logger.mu.Unlock()
	ls.Stack = make([]*LogMsg, 0)
}

func (ls *LogStack) AddAndPrint(lvl LogLevel, msg string) {
	ls.Add(lvl, msg)
	ls.PrintStack()
}

func NewLogStack(l *Logger) *LogStack {
	return &LogStack{
		Logger: l,
		Stack:  make([]*LogMsg, 0),
	}
}

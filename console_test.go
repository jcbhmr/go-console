package console

import "testing"

func ptr[T any](v T) *T {
	return &v
}

func TestConsoleAssert(t *testing.T) {
	ConsoleAssert(ptr(false), "This is an error message")
}

func TestConsoleClear(t *testing.T) {
	ConsoleClear()
}

func TestConsoleDebug(t *testing.T) {
	ConsoleDebug("This is a debug message")
}

func TestConsoleError(t *testing.T) {
	ConsoleError("This is an error message")
}

func TestConsoleInfo(t *testing.T) {
	ConsoleInfo("This is an info message")
}

func TestConsoleLog(t *testing.T) {
	ConsoleLog("This is a log message")
}

func TestConsoleTable(t *testing.T) {
	ConsoleTable([]struct{ A, B int }{{1, 2}, {3, 4}}, []string{"A", "B"})
}

func TestConsoleTrace(t *testing.T) {
	ConsoleTrace("This is a trace message")
}

func TestConsoleWarn(t *testing.T) {
	ConsoleWarn("This is a warning message")
}

func TestConsoleDir(t *testing.T) {
	ConsoleDir(1, nil)
}

func TestConsoleDirxml(t *testing.T) {
	ConsoleDirxml(1)
}

func TestConsoleCount(t *testing.T) {
	ConsoleCount(ptr("label"))
}

func TestConsoleGroup(t *testing.T) {
	ConsoleGroup("label")
	ConsoleGroupEnd()
}

func TestConsoleGroupCollapsed(t *testing.T) {
	ConsoleGroupCollapsed("label")
	ConsoleGroupEnd()
}

func TestConsoleTime(t *testing.T) {
	ConsoleTime(ptr("label"))
	ConsoleTimeEnd(ptr("label"))
}

func TestConsoleTimeLog(t *testing.T) {
	ConsoleTimeLog(ptr("label"), "This is a time log message")
}

func TestConsoleTimeEnd(t *testing.T) {
	ConsoleTimeEnd(ptr("label"))
}
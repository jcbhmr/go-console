package console

import "syscall/js"

func ConsoleAssert(condition *bool, data ...any) {
	if condition == nil && len(data) == 0 {
		js.Global().Get("console").Call("assert")
	} else {
		args := append([]any{*condition}, data...)
		js.Global().Get("console").Call("assert", args)
	}
}

func ConsoleClear() {
	js.Global().Get("console").Call("clear")
}

func ConsoleDebug(data ...any) {
	js.Global().Get("console").Call("debug", data...)
}

func ConsoleError(data ...any) {
	js.Global().Get("console").Call("error", data...)
}

func ConsoleInfo(data ...any) {
	js.Global().Get("console").Call("info", data...)
}

func ConsoleLog(data ...any) {
	js.Global().Get("console").Call("log", data...)
}

func ConsoleTable(tabularData any, properties []string) {
	if tabularData == nil && properties == nil {
		js.Global().Get("console").Call("table")
	} else if properties == nil {
		js.Global().Get("console").Call("table", tabularData)
	} else {
		js.Global().Get("console").Call("table", tabularData, properties)
	}
}

func ConsoleTrace(data ...any) {
	js.Global().Get("console").Call("trace", data...)
}

func ConsoleWarn(data ...any) {
	js.Global().Get("console").Call("warn", data...)
}

func ConsoleDir(item any, options any) {
	if options == nil {
		js.Global().Get("console").Call("dir", item)
	} else {
		js.Global().Get("console").Call("dir", item, options)
	}
}

func ConsoleDirxml(data ...any) {
	js.Global().Get("console").Call("dirxml", data...)
}

func ConsoleCount(label *string) {
	if label == nil {
		js.Global().Get("console").Call("count")
	} else {
		js.Global().Get("console").Call("count", *label)
	}
}

func ConsoleCountReset(label *string) {
	if label == nil {
		js.Global().Get("console").Call("countReset")
	} else {
		js.Global().Get("console").Call("countReset", *label)
	}
}

func ConsoleGroup(data ...any) {
	js.Global().Get("console").Call("group", data...)
}

func ConsoleGroupCollapsed(data ...any) {
	js.Global().Get("console").Call("groupCollapsed", data...)
}

func ConsoleGroupEnd() {
	js.Global().Get("console").Call("groupEnd")
}

func ConsoleTime(label *string) {
	if label == nil {
		js.Global().Get("console").Call("time")
	} else {
		js.Global().Get("console").Call("time", *label)
	}
}

func ConsoleTimeLog(label *string, data ...any) {
	if label == nil && len(data) == 0 {
		js.Global().Get("console").Call("timeLog")
	} else {
		args := append([]any{*label}, data...)
		js.Global().Get("console").Call("timeLog", args)
	}
}

func ConsoleTimeEnd(label *string) {
	if label == nil {
		js.Global().Get("console").Call("timeEnd")
	} else {
		js.Global().Get("console").Call("timeEnd", *label)
	}
}

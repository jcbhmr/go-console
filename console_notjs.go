//go:build !js

package console

import (
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"golang.org/x/term"
)

func logger(logLevel string, args []any) {
	// 1. If args is empty, return.
	if len(args) == 0 {
		return
	}

	// 2. Let first be args[0].
	first := args[0]

	// 3. Let rest be all elements following first in args.
	rest := args[1:]

	// 4. If rest is empty, perform Printer(logLevel, « first ») and return.
	if len(rest) == 0 {
		printer(logLevel, []any{first}, nil)
		return
	} else {
		// 5. Otherwise, perform Printer(logLevel, Formatter(args)).
		printer(logLevel, formatter(args), nil)
	}

	// 6. Return undefined.
}

func formatter(args []any) []any {
	// 1. If args’s size is 1, return args.
	if len(args) == 1 {
		return args
	}

	// 2. Let target be the first element of args.
	target := args[0].(string)

	// 3. Let current be the second element of args.
	current := args[1]

	// 4. Find the first possible format specifier specifier, from the left to the right in target.
	match := regexp.MustCompile(`%[sdifoOc]`).FindStringIndex(target)
	var specifier string
	if match != nil {
		specifier = target[match[0]:match[1]]
	}

	// 5. If no format specifier was found, return args.
	if specifier == "" {
		return args
	} else {
		// 6. Otherwise:
		var converted any

		// 1. If specifier is %s, let converted be the result of Call(%String%, undefined, « current »).
		if specifier == "%s" {
			converted = fmt.Sprintf("%s", current)
		}

		// 2. If specifier is %d or %i:
		if specifier == "%d" || specifier == "%i" {

			// 1. If current is a Symbol, let converted be NaN
			if false {
				converted = "NaN"
			} else {
				// 2. Otherwise, let converted be the result of Call(%parseInt%, undefined, « current, 10 »).
				converted2, err := strconv.ParseInt(fmt.Sprintf("%s", current), 10, 32)
				if err != nil {
					converted = "NaN"
				} else {
					converted = converted2
				}
			}
		}

		// 3. If specifier is %f:
		if specifier == "%f" {

			// 1. If current is a Symbol, let converted be NaN
			if false {
				converted = "NaN"
			} else {
				// 2. Otherwise, let converted be the result of Call(%parseFloat%, undefined, « current »).
				converted2, err := strconv.ParseFloat(fmt.Sprintf("%s", current), 64)
				if err != nil {
					converted = "NaN"
				} else {
					converted = converted2
				}
			}
		}

		// 4. If specifier is %o, optionally let converted be current with optimally useful formatting applied.
		if specifier == "%o" {
			// converted = fmt.Sprintf("%#+v", current)
		}

		// 5. If specifier is %O, optionally let converted be current with generic JavaScript object formatting applied.
		if specifier == "%O" {
			// converted = fmt.Sprintf("%#v", current)
		}

		// 6. TODO: process %c
		if specifier == "%c" {
			// Ignore for now.
			converted = ""
		}

		// 7. If any of the previous steps set converted, replace specifier in target with converted.
		if converted != nil {
			target = target[:match[0]] + fmt.Sprint(converted) + target[match[1]:]
		}
	}

	// 7. Let result be a list containing target together with the elements of args starting from the third onward.
	result := append([]any{target}, args[2:]...)

	// 8. Return Formatter(result).
	return formatter(result)
}

func printer(logLevel string, args []any, options any) {
	argsStrings := make([]string, len(args))
	for i, arg := range args {
		argsStrings[i] = fmt.Sprint(arg)
	}
	joined := strings.Join(argsStrings, " ")
	lines := strings.Split(joined, "\n")
	indented := []string{}
	for _, line := range lines {
		for range groupStack {
			line = "  " + line
		}
		indented = append(indented, line)
	}
	message := strings.Join(indented, "\n")
	switch logLevel {
	case "assert":
		fallthrough
	case "error":
		fallthrough
	case "warn":
		fmt.Fprintln(os.Stderr, message)
	default:
		fmt.Println(message)
	}
	_ = options
}

func ConsoleAssert(condition *bool, data ...any) {
	var condition2 bool
	if condition == nil {
		condition2 = false
	} else {
		condition2 = *condition
	}

	// 1. If condition is true, return.
	if condition2 {
		return
	}

	// 2. Let message be a string without any formatting specifiers indicating generically an assertion failure (such as "Assertion failed").
	message := "Assertion failed"

	// 3. If data is empty, append message to data.
	if len(data) == 0 {
		data = append(data, message)
	} else {
		// 4. Otherwise:

		// 1. Let first be data[0].
		first := data[0]

		// 2. If first is not a String, then prepend message to data.
		if firstString, ok := first.(string); !ok {
			data = append([]any{message}, data...)
		} else {
			// 3. Otherwise:

			// 1. Let concat be the concatenation of message, U+003A (:), U+0020 SPACE, and first.
			concat := message + ": " + firstString

			// 2. Set data[0] to concat.
			data[0] = concat
		}
	}

	// 5. Perform Logger("assert", data).
	logger("assert", data)
}

func ConsoleClear() {
	// 1. Empty the appropriate group stack.
	groupStack = []group{}

	// 2. If possible for the environment, clear the console. (Otherwise, do nothing.)
	if term.IsTerminal(int(os.Stdout.Fd())) {
		fmt.Print("\033[H\033[2J")
	}
}

func ConsoleDebug(data ...any) {
	// 1. Perform Logger("debug", data).
	logger("debug", data)
}

func ConsoleError(data ...any) {
	// 1. Perform Logger("error", data).
	logger("error", data)
}

func ConsoleInfo(data ...any) {
	// 1. Perform Logger("info", data).
	logger("info", data)
}

func ConsoleLog(data ...any) {
	// 1. Perform Logger("log", data).
	logger("log", data)
}

func ConsoleTable(tabularData any, properties []string) {
	logger("log", []any{tabularData})
}

func ConsoleTrace(data ...any) {
	// 1. Let trace be some implementation-defined, potentially-interactive representation of the callstack from where this function was called.
	var trace string
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		trace = ""
	} else {
		runtimeFunc := runtime.FuncForPC(pc)
		trace = fmt.Sprintf("%s (%s:%d)", runtimeFunc.Name(), file, line)
	}

	// 2. Optionally, let formattedData be the result of Formatter(data), and incorporate formattedData as a label for trace.
	// TODO

	// 3. Perform Printer("trace", « trace »).
	printer("trace", []any{trace}, nil)
}

func ConsoleWarn(data ...any) {
	// 1. Perform Logger("warn", data).
	logger("warn", data)
}

func ConsoleDir(item any, options any) {
	// 1. Let object be item with generic JavaScript object formatting applied.
	object := fmt.Sprintf("%#v", item)

	// 2. Perform Printer("dir", « object », options).
	printer("dir", []any{object}, options)
}

func ConsoleDirxml(data ...any) {
	// 1. Let finalList be a new list, initially empty.
	finalList := []any{}

	// 2. For each item of data:
	for _, item := range data {
		// 1. Let converted be a DOM tree representation of item if possible; otherwise let converted be item with optimally useful formatting applied.
		converted := fmt.Sprintf("%#+v", item)

		// 2. Append converted to finalList.
		finalList = append(finalList, converted)
	}

	// 3. Perform Logger("dirxml", finalList).
	logger("dirxml", finalList)
}

var countMap = map[string]int{}

func ConsoleCount(label *string) {
	var label2 string
	if label == nil {
		label2 = "default"
	} else {
		label2 = *label
	}

	// 1. Let map be the associated count map.
	map2 := countMap

	// 2. If map[label] exists, set map[label] to map[label] + 1.
	if _, ok := map2[label2]; ok {
		map2[label2] = map2[label2] + 1
	} else {
		// 3. Otherwise, set map[label] to 1.
		map2[label2] = 1
	}

	// 4. Let concat be the concatenation of label, U+003A (:), U+0020 SPACE, and ToString(map[label]).
	concat := label2 + ": " + fmt.Sprint(map2[label2])

	// 5. Perform Logger("count", « concat »).
	logger("count", []any{concat})
}

func ConsoleCountReset(label *string) {
	var label2 string
	if label == nil {
		label2 = "default"
	} else {
		label2 = *label
	}

	// 1. Let map be the associated count map.
	map2 := countMap

	// 2. If map[label] exists, set map[label] to 0.
	if _, ok := map2[label2]; ok {
		map2[label2] = 0
	} else {
		// 3. Otherwise:

		// 1. Let message be a string without any formatting specifiers indicating generically that the given label does not have an associated count.
		message := "The given label does not have an associated count"

		// 2. Perform Logger("countReset", « message »);
		logger("countReset", []any{message})
	}
}

type group struct {
	label []any
}

func (g group) String() string {
	labelStrings := make([]string, len(g.label))
	for i, label := range g.label {
		labelStrings[i] = fmt.Sprint(label)
	}
	return strings.Join(labelStrings, " ")
}

var groupStack = []group{}

func ConsoleGroup(data ...any) {
	// 1. Let group be a new group.
	group2 := group{}

	// 2. If data is not empty, let groupLabel be the result of Formatter(data). Otherwise, let groupLabel be an implementation-chosen label representing a group.
	var groupLabel []any
	if len(data) != 0 {
		groupLabel = formatter(data)
	} else {
		groupLabel = []any{""}
	}

	// 3. Incorporate groupLabel as a label for group.
	group2.label = groupLabel

	// 4. Optionally, if the environment supports interactive groups, group should be expanded by default.

	// 5. Perform Printer("group", « group »).
	printer("group", []any{group2}, nil)

	// 6. Push group onto the appropriate group stack.
	groupStack = append(groupStack, group2)
}

func ConsoleGroupCollapsed(data ...any) {
	// 1. Let group be a new group.
	group2 := group{}

	// 2. If data is not empty, let groupLabel be the result of Formatter(data). Otherwise, let groupLabel be an implementation-chosen label representing a group.
	var groupLabel []any
	if len(data) != 0 {
		groupLabel = formatter(data)
	} else {
		groupLabel = []any{""}
	}

	// 3. Incorporate groupLabel as a label for group.
	group2.label = groupLabel

	// 4. Optionally, if the environment supports interactive groups, group should be collapsed by default.

	// 5. Perform Printer("groupCollapsed", « group »).
	printer("groupCollapsed", []any{group2}, nil)

	// 6. Push group onto the appropriate group stack.
	groupStack = append(groupStack, group2)
}

func ConsoleGroupEnd() {
	// 1. Pop the last group from the group stack.
	_ = groupStack[len(groupStack)-1]
	groupStack = groupStack[:len(groupStack)-1]
}

var timerTable = map[string]time.Time{}

func ConsoleTime(label *string) {
	var label2 string
	if label == nil {
		label2 = "default"
	} else {
		label2 = *label
	}

	// 1. If the associated timer table contains an entry with key label, return, optionally reporting a warning to the console indicating that a timer with label label has already been started.
	if _, ok := timerTable[label2]; ok {
		ConsoleWarn("A timer with label " + label2 + " has already been started")
		return
	} else {
		// 2. Otherwise, set the value of the entry with key label in the associated timer table to the current time.
		timerTable[label2] = time.Now()
	}
}

func ConsoleTimeLog(label *string, data ...any) {
	var label2 string
	if label == nil {
		label2 = "default"
	} else {
		label2 = *label
	}

	// 1. Let timerTable be the associated timer table.
	timerTable2 := timerTable

	// 2. Let startTime be timerTable[label].
	startTime, ok := timerTable2[label2]
	if !ok {
		ConsoleWarn("A timer with label " + label2 + " has not been started")
		return
	}

	// 3. Let duration be a string representing the difference between the current time and startTime, in an implementation-defined format.
	duration := time.Since(startTime).String()

	// 4. Let concat be the concatenation of label, U+003A (:), U+0020 SPACE, and duration.
	concat := label2 + ": " + duration

	// 5. Prepend concat to data.
	data = append([]any{concat}, data...)

	// 6. Perform Printer("timeLog", data).
	printer("timeLog", data, nil)
}

func ConsoleTimeEnd(label *string) {
	var label2 string
	if label == nil {
		label2 = "default"
	} else {
		label2 = *label
	}

	// 1. Let timerTable be the associated timer table.
	timerTable2 := timerTable

	// 2. Let startTime be timerTable[label].
	startTime, ok := timerTable2[label2]
	if !ok {
		ConsoleWarn("A timer with label " + label2 + " has not been started")
		return
	}

	// 3. Remove timerTable[label].
	delete(timerTable2, label2)

	// 4. Let duration be a string representing the difference between the current time and startTime, in an implementation-defined format.
	duration := time.Since(startTime).String()

	// 5. Let concat be the concatenation of label, U+003A (:), U+0020 SPACE, and duration.
	concat := label2 + ": " + duration

	// 6. Perform Printer("timeEnd", « concat »).
	printer("timeEnd", []any{concat}, nil)
}

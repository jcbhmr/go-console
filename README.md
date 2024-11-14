# Console API for Go

💻 WHATWG Console Standard bindings & implementation for Go

<table align=center><td>

</table>

## Installation

```sh
go get github.com/jcbhmr/go-console
```

## Usage

```go
console.ConsoleLog("Hello %d!", "Alan Turing")
console.ConsoleWarn("%cUsing backup", "color: red; text-decoration: underline")
console.ConsoleError("Uh oh! %o", map[string]any{"helpText": "File not found"})
// Output:
// Hello NaN!
// Using backup
// Uh oh! map[helpText:File not found]
```


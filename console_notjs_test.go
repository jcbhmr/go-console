//go:build !js

package console

import "testing"

func TestPrinter(t *testing.T) {
	printer("log", []any{"Hi", 1, 2, 3, "four", []int{5, 6, 7}, map[string]any{"eight": 8}}, nil)
}

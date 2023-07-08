package graph

import "fmt"

func PrintIf(condition bool, format string, args ...any) {
	if condition {
		fmt.Printf(format, args...)
	}
}

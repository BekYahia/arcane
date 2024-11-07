package parser

/**
* Usage:
	defer unTrace(trace("myFunction"))

A deferred function's arguments are evaluated when the defer statement is evaluated.
and the deferred function calls are executed in Last in First out order.

deferred function: unTrace, and the arguments is trace.
*/

import (
	"fmt"
	"strings"
)

var traceLevel = 0 // starts with -1 so the first traced function has zero indentation/
const traceIndentPlaceholder = "\t"

func indentLevel() string {
	return strings.Repeat(traceIndentPlaceholder, traceLevel-1)
}
func printTrace(msg string) {
	fmt.Printf("%s%s\n", indentLevel(), msg)
}
func incrementIndent() { traceLevel += 1 }
func decrementIndent() { traceLevel -= 1 }

func trace(msg string) string {
	incrementIndent()
	printTrace("BEGIN: " + msg)
	return msg
}
func unTrace(msg string) {
	printTrace("END: " + msg)
	decrementIndent()
}

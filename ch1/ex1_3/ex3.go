package ex3

import "strings"

// JoinByFor ...
func JoinByFor(args []string) string {
	var s, sep string
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	return s
}

// JoinByJoin ...
func JoinByJoin(args []string) string {
	return strings.Join(args, " ")
}

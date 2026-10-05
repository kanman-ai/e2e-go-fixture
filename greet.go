// Package greet is a minimal Go fixture for kanman's end-to-end checks of
// team environments (a Go toolchain only inside the team image).
package greet

import "fmt"

// Hello returns a greeting for name.
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

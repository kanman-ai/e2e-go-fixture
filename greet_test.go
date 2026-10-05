package greet

import "testing"

func TestHello(t *testing.T) {
	if got := Hello("kanman"); got != "Hello, kanman" {
		t.Fatalf("Hello() = %q", got)
	}
}

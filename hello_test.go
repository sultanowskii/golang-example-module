package example

import "testing"

func TestHello(t *testing.T) {
    want := "howdy!"
    if got := Hello(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}

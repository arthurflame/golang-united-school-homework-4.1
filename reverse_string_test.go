package reverse_string

import "testing"

func TestReverseString(t *testing.T) {
	got := ReverseString("Привет")
	want := "тевирП"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

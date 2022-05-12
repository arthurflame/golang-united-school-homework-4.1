package reverse_string

import "testing"

func TestReverseString(t *testing.T) {
	got := ReverseString("Hello world!")
	want := "!dlrow olleH"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

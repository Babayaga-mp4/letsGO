package hello

import (
	"testing"
)

func TestSayHello(t *testing.T) {
	want := "Hello, Test!"
	got := Say([]string{"Test"})

	if want != got {
		t.Errorf("wanted %s, got %s", want, got)
	}
}

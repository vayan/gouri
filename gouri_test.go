package gouri

import (
	//"bitbucket.org/vayan/gouri"
	"testing"
)

func TestNew(t *testing.T) {
	s := New(5)
	if len(s) != 5 {
		t.Fatal("wrong length")
	}
}

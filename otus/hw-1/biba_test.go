package main

import (
	"testing"
)

func TestFirstOne(t *testing.T) {
	in := "a4bc2d5e"
	er := "aaaabccddddde"
	if ar := First(in); ar != er {
		t.Fatalf("bad unarchive for %s: \nexpected result %s \nactual result   %s", in, er, ar)
	}
}

func TestFirstTwo(t *testing.T) {
	in := "abcd"
	er := "abcd"
	if ar := First(in); ar != er {
		t.Fatalf("bad unarchive for %s: \nexpected result %s \nactual result   %s", in, er, ar)
	}
}

func TestFirstThree(t *testing.T) {
	in := "45"
	er := ""
	if ar := First(in); ar != er {
		t.Fatalf("bad unarchive for %s: \nexpected result %s \nactual result   %s", in, er, ar)
	}
}

func TestFirstFour(t *testing.T) {
	in := ""
	er := ""
	if ar := First(in); ar != er {
		t.Fatalf("bad unarchive for %s: \nexpected result %s \nactual result   %s", in, er, ar)
	}
}

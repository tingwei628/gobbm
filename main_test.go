package main

import "testing"

func TestMain(t *testing.T) {
	got := 1
	want := 1
	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}

package main

import "testing"

func TestHello(t *testing.T)  {
	got := Hello()
	want := "hello world2"

	t.Log("doing")

	if got != want {
		t.Errorf("expect '%s' but '%q' ", want, got)
	}
}
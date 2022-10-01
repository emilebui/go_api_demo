package demo

import "testing"

func TestHelloWorldLogic(t *testing.T) {
	msg := HelloWorldLogic()
	if msg != "Hello World" {
		t.Fatalf("Fail in logic")
	}
}

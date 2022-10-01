package demo

import (
	"context"
	"go_api/proto_gen"
	"testing"
)

func TestServer_HelloWorld(t *testing.T) {
	server := &Server{}
	r, err := server.HelloWorld(context.Background(), &proto_gen.EmptyMessage{})
	if err != nil {
		t.Fatalf("something is wrong, %v", err)
	}
	if r.Msg != "Hello World" {
		t.Fatalf("Output is different")
	}
}

package main

import (
	"testing"

	"github.com/rotemtam/ent-grpc-example/ent/proto/entpb"
)

func TestUserProto(t *testing.T) {
	user := entpb.User{
		Name:         "rotemtam",
		EmailAddress: "rotemtam@example.com",
	}
	if user.GetName() != "rotemtam" {
		t.Fatal("expected user name to be rotemtam")
	}
	if user.GetEmailAddress() != "rotemtam@example.com" {
		t.Fatal("expected email address to be rotemtam@example.com")
	}
}


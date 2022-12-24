package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/rotemtam/ent-grpc-example/ent/proto/entpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Open a connection to the server.
	conn, err := grpc.Dial(":5000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed connecting to server: %s", err)
	}
	defer conn.Close()

	// Create a User service Client on the connection.
	client := entpb.NewUserServiceClient(conn)

	// Ask the server to create a random User.
	ctx := context.Background()
	user := randomUser()
	created, err := client.Create(ctx, &entpb.CreateUserRequest{
		User: user,
	})
	if err != nil {
		se, _ := status.FromError(err)
		log.Fatalf("failed creating user: status=%s message=%s", se.Code(), se.Message())
	}
	log.Printf("user created with id: %d", created.Id)

	// On a separate RPC invocation, retrieve the user we saved previously.
	get, err := client.Get(ctx, &entpb.GetUserRequest{
		Id: created.Id,
	})
	if err != nil {
		se, _ := status.FromError(err)
		log.Fatalf("failed retrieving user: status=%s message=%s", se.Code(), se.Message())
	}
	log.Printf("retrieved user with id=%d: %v", get.Id, get)

	grpClient := entpb.NewGroupServiceClient(conn)
	g, err := grpClient.Create(ctx, &entpb.CreateGroupRequest{
		Group: &entpb.Group{
			Name: "test",
			Admin: &entpb.User{
				Id: created.Id,
			},
		},
	})
	if err != nil {
		se, _ := status.FromError(err)
		log.Fatalf("failed creating group: status=%s message=%s", se.Code(), se.Message())
	}
	log.Printf("group created with id: %d", g.Id)
	g, err = grpClient.Get(ctx, &entpb.GetGroupRequest{
		Id:   g.Id,
		View: entpb.GetGroupRequest_WITH_EDGE_IDS,
	})
	if err != nil {
		se, _ := status.FromError(err)
		log.Fatalf("failed creating group: status=%s message=%s", se.Code(), se.Message())
	}
	log.Printf("group retrieved with id: %d admin id: %d", g.Id, g.Admin.Id)

}

func randomUser() *entpb.User {
	r := rand.Int()
	return &entpb.User{
		Username:  fmt.Sprintf("user_%d", r),
		FirstName: fmt.Sprintf("first_%d", r),
		LastName:  fmt.Sprintf("last_%d", r),
		Email:     fmt.Sprintf("user_%d@example.com", r),
	}
}

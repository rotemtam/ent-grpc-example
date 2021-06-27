package main

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/rotemtam/ent-grpc-example/ent/category"
	"github.com/rotemtam/ent-grpc-example/ent/enttest"
	"github.com/rotemtam/ent-grpc-example/ent/proto/entpb"
	"github.com/rotemtam/ent-grpc-example/ent/user"
)

func TestServiceWithEdges(t *testing.T) {
	// start by initializing an ent client connected to an in memory sqlite instance
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	// next, initialize the UserService. Notice we won't be opening an actual port and creating a gRPC server
	// and instead we are just calling the library code directly. 
	svc := entpb.NewUserService(client)

	// next, we create a category directly using the ent client. notice we are initializing it with no relation
	// to a User.
	cat := client.Category.Create().SetName("cat_1").SaveX(ctx)

	// next, we invoke the User service's `Create` method. Notice we are passing a list of entpb.Category 
	// instances with only the ID set. 
	create, err := svc.Create(ctx, &entpb.CreateUserRequest{
		User: &entpb.User{
			Name:         "user",
			EmailAddress: "user@service.code",
			Administered: []*entpb.Category{
				{Id: int32(cat.ID)},
			},
		},
	})
	if err != nil {
		t.Fatal("failed creating user using UserService", err)
	}

	// to verify everything worked correctly, we query the category table to check we have exactly
	// one category which is administered by the created user.
	count, err := client.Category.
		Query().
		Where(
			category.HasAdminWith(
				user.ID(int(create.Id)),
			),
		).
		Count(ctx)
	if err != nil {
		t.Fatal("failed counting categories admin by created user", err)
	}
	if count != 1 {
		t.Fatal("expected exactly one group to managed by the created user")
	}
}

func TestGet(t *testing.T) {
	// start by initializing an ent client connected to an in memory sqlite instance
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	// next, initialize the UserService. Notice we won't be opening an actual port and creating a gRPC server
	// and instead we are just calling the library code directly.
	svc := entpb.NewUserService(client)

	// next, create a user, a category and set that user to be the admin of the category
	user := client.User.Create().
		SetName("rotemtam").
		SetEmailAddress("r@entgo.io").
		SaveX(ctx)

	client.Category.Create().
		SetName("category").
		SetAdmin(user).
		SaveX(ctx)

	// next, retrieve the user without edge information
	get, err := svc.Get(ctx, &entpb.GetUserRequest{
		Id: int32(user.ID),
	})
	if err != nil {
		t.Fatal("failed retrieving the created user", err)
	}
	if len(get.Administered) != 0 {
		t.Fatal("by default edge information is not supposed to be retrieved")
	}

	// next, retrieve the user *WITH* edge information
	get, err = svc.Get(ctx, &entpb.GetUserRequest{
		Id:   int32(user.ID),
		View: entpb.GetUserRequest_WITH_EDGE_IDS,
	})
	if err != nil {
		t.Fatal("failed retrieving the created user", err)
	}
	if len(get.Administered) != 1 {
		t.Fatal("using WITH_EDGE_IDS edges should be returned")
	}
}

package main

import (
	"context"
	"fmt"
	pb "github.com/ergildo/go-lang-grcp-app-mysql-example/user-pb"
	"google.golang.org/grpc"
	"log"
	"time"
)

const address = "localhost:50051"

func main() {
	dial, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	defer dial.Close()
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := pb.NewUserServiceBPClient(dial)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	createAllUsers(c, ctx)

	users := listAllUser(c, ctx)

	printUsers(users)
	findAllUserById(c, ctx, users)
	updateAllUsers(c, ctx, users)
	deleteAllUsers(c, ctx, users)

}

func printUsers(users []*pb.UserResponse) {
	for _, user := range users {
		printUser(user)
	}
}

func deleteAllUsers(c pb.UserServiceBPClient, ctx context.Context, users []*pb.UserResponse) {
	log.Println("delete all users...")
	for _, user := range users {
		deleteUser(c, ctx, user.GetId())
	}
}

func updateAllUsers(c pb.UserServiceBPClient, ctx context.Context, users []*pb.UserResponse) {
	log.Println("updating users...")
	for i, user := range users {
		newName := fmt.Sprintf("user_updated_%d", i)
		u := updateUser(c, ctx, &pb.UserResponse{Id: user.GetId(), Name: newName})
		printUser(u)
	}
}

func findAllUserById(c pb.UserServiceBPClient, ctx context.Context, users []*pb.UserResponse) {
	log.Println("finding users by id...")
	for _, user := range users {
		u := findUserById(c, ctx, user.GetId())
		printUser(u)
	}
}

func deleteUser(c pb.UserServiceBPClient, ctx context.Context, id int64) {
	_, err := c.DeleteUser(ctx, &pb.UserRequest{
		Id: id,
	})
	if err != nil {
		log.Fatalf("could not list user by id: %v", err)
	}
}

func updateUser(c pb.UserServiceBPClient, ctx context.Context, user *pb.UserResponse) *pb.UserResponse {
	u, err := c.UpdateUser(ctx, &pb.UpdateUserRequest{Id: user.GetId(), Name: user.GetName()})
	if err != nil {
		log.Fatalf("could not list user by id: %v", err)
	}

	return u

}

func findUserById(c pb.UserServiceBPClient, ctx context.Context, id int64) *pb.UserResponse {
	user, err := c.FindUserById(ctx, &pb.UserRequest{Id: id})
	if err != nil {
		log.Fatalf("could not list user by id: %v", err)
	}

	return user

}

func listAllUser(c pb.UserServiceBPClient, ctx context.Context) []*pb.UserResponse {
	log.Println("listing all users...")
	ru, err := c.ListAllUsers(ctx, &pb.Void{})
	if err != nil {
		log.Fatalf("could not list users: %v", err)
	}

	return ru.GetUserResponse()

}

func createAllUsers(c pb.UserServiceBPClient, ctx context.Context) {
	log.Println("crating all users...")
	var users []*pb.UserResponse
	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("user_%d", i)
		userResponse := createNewUser(c, ctx, &pb.NewUserRequest{
			Name: name,
		})
		users = append(users, userResponse)
	}

}

func createNewUser(c pb.UserServiceBPClient, ctx context.Context, newUser *pb.NewUserRequest) *pb.UserResponse {

	userResponse, err := c.CreateUser(ctx, newUser)

	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}
	return userResponse
}

func printUser(userResponse *pb.UserResponse) (int, error) {
	return fmt.Printf("user{id:%d, name:%s}\n", userResponse.GetId(), userResponse.GetName())
}

package main

import (
	"context"
	pb "github.com/ergildo/go-lang-grcp-app-mysql-example/user-pb"
	"github.com/ergildo/go-lang-grcp-app-mysql-example/user-server/model"
	"github.com/ergildo/go-lang-grcp-app-mysql-example/user-server/service"
	"github.com/ergildo/go-lang-grcp-app-mysql-example/user-server/setup"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = ":50051"

type server struct {
	pb.UnimplementedUserServiceBPServer
}

func (s *server) CreateUser(ctx context.Context, in *pb.NewUserRequest) (*pb.UserResponse, error) {
	user := service.Save(model.User{Name: in.GetName()})
	return &pb.UserResponse{Id: user.Id, Name: user.Name}, nil
}
func (s *server) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	user := service.Update(model.User{Id: in.Id, Name: in.Name})
	return &pb.UserResponse{
		Id: user.Id, Name: user.Name,
	}, nil
}
func (s *server) FindUserById(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	user := service.FindById(in.Id)
	return &pb.UserResponse{
		Id: user.Id, Name: user.Name,
	}, nil
}
func (s *server) ListAllUsers(ctx context.Context, in *pb.Void) (*pb.ListAllUsersResponse, error) {
	users := service.ListAll()
	var userResponses []*pb.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, &pb.UserResponse{
			Id: user.Id, Name: user.Name,
		})
	}
	return &pb.ListAllUsersResponse{UserResponse: userResponses}, nil
}
func (s *server) DeleteUser(ctx context.Context, in *pb.UserRequest) (*pb.Void, error) {
	service.Delete(in.Id)
	return &pb.Void{}, nil
}

func main() {
	setup.SetUpDB()
	startGrpcServer()
}

func startGrpcServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceBPServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

package user

import (
	"context"
	"strconv"

	pb "github.com/guutong/demo-gin/proto"
)

type UserServiceServer struct {
	repo IUserRepository
	pb.UnimplementedUserServiceServer
}

func NewUserServiceServer(repo IUserRepository) *UserServiceServer {
	return &UserServiceServer{repo: repo}
}

func (service *UserServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	newUser := User{Name: req.Name}

	err := service.repo.NewUser(&newUser)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{Data: "User created successfully!"}, nil
}

func (service *UserServiceServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	newUser := User{Name: req.Name}
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, err
	}

	err = service.repo.UpdateUser(id, &newUser)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateUserResponse{Data: "User updated successfully!"}, nil
}

func (service *UserServiceServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, err
	}

	err = service.repo.DeleteUser(id)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteUserResponse{Data: "User details deleted successfully!"}, nil
}

func (service *UserServiceServer) GetAllUsers(context.Context, *pb.Empty) (*pb.GetAllUsersResponse, error) {
	resp, err := service.repo.GetUser()
	var users []*pb.UserResponse

	if err != nil {
		return nil, err
	}

	for _, v := range resp {
		var singleUser = &pb.UserResponse{
			Id:   strconv.Itoa(int(v.ID)),
			Name: v.Name,
		}
		users = append(users, singleUser)
	}

	return &pb.GetAllUsersResponse{Users: users}, nil
}

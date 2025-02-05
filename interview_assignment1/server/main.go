package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/shubhammurarka/grpc/proto"
	"github.com/shubhammurarka/grpc/server/AuthUser"
	"github.com/shubhammurarka/grpc/server/UserStore"
	"google.golang.org/grpc"
)

var (
	secretKey = "SECRETKEY"
	duration  = 20 * time.Second
)

type AuthService struct {
	pb.UnimplementedAuthServiceServer
	user UserStore.UserStoreInterface
	auth AuthUser.JwtManagerInterface
}

func NewAuthService(user UserStore.UserStoreInterface, auth AuthUser.JwtManagerInterface) *AuthService {
	return &AuthService{
		user: user,
		auth: auth,
	}
}

func (a *AuthService) Login(ctx context.Context, rq *pb.LoginRequest) (*pb.LoginResponse, error) {
	user := a.user.Find(rq.UserName)
	if user == nil {
		return nil, fmt.Errorf("No User Found")
	}

	is := a.user.IsCorrectPassword(user.UserName, rq.Password)
	if !is {
		return nil, fmt.Errorf("Not Authorised")
	}

	token, err := a.auth.GenerateToken(user.UserName)
	if err != nil {
		return nil, fmt.Errorf("Internal Server Error")
	}

	res := &pb.LoginResponse{
		AccessToken: token,
	}

	return res, nil
}

func (a *AuthService) Register(ctx context.Context, rq *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user, err := a.user.Save(rq.GetUserName(), rq.GetPassword())
	if err != nil {
		return nil, err
	}

	saveresponse := fmt.Sprintf("User Succesfully saved")
	res := &pb.RegisterResponse{
		UserName: user.UserName,
		Response: saveresponse,
	}

	// for i, val := a.

	return res, nil
}

func (a *AuthService) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {

	_, err := a.auth.VerifyToken(req.GetAccessToken())
	if err != nil {
		return nil, fmt.Errorf("Invalid token")
	}

	err = a.auth.Logout(req.GetAccessToken())

	if err != nil {
		return nil, fmt.Errorf("error logging out")
	}

	res := &pb.LogoutResponse{
		Response: fmt.Sprintf("User Successfully logged out"),
	}

	return res, nil
}

func main() {
	list, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Printf("error listening on port 8080")
	}

	auth := AuthUser.NewJwtManager(secretKey, duration)
	user := UserStore.NewUserStore()

	authserver := NewAuthService(user, auth)

	server := grpc.NewServer()
	pb.RegisterAuthServiceServer(server, authserver)
	server.Serve(list)
}

package grpc

import (
	"context"
	userspb "reg/contracts/users"
	"reg/users-service/internal/domain"
	"reg/users-service/internal/usecase"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Config struct {
	UcCreateUser *usecase.CreateUC

	UcDeleteUser *usecase.DeleteUC

	// UcCheckAuth *usecase.CheckAuthUC

	UcGetUserByLogin *usecase.GetUserByLoginUC
}

type Server struct {
	userspb.UnimplementedUserServiceServer
	config Config
}

func NewServer(conf Config) *Server {
	return &Server{
		config: conf,
	}
}

func (s *Server) CreateUser(
	ctx context.Context,
	req *userspb.CreateUserRequest,
) (*userspb.CreateUserResponse, error) {
	userID, err := s.config.UcCreateUser.CreateUser(
		ctx,
		&domain.User{
			Login: req.GetLogin(),
			PasswordHash: req.GetPasswordHash(),
		},
	)

	if err != nil {
		return nil, err
	}

	return &userspb.CreateUserResponse{
		UserId: userID,
	}, nil
}

func (s *Server) DeleteUser(
	ctx context.Context,
	req *userspb.DeleteUserRequest,
) (*emptypb.Empty, error) {
	err := s.config.UcDeleteUser.DeleteUser(ctx, req.GetUserId())

	return &emptypb.Empty{}, err
}

// func (s *Server) CheckAuth(
// 	ctx context.Context,
// 	req *userspb.CheckAuthRequest,
// )  (userspb.CheckAuthResponse, error) {
// 	res, err := s.config.UcCheckAuth.CheckAuth(
// 		ctx,
// 		&domain.User{
// 			Login: req.GetLogin(),
// 			PasswordHash: req.GetPasswordHash(),
// 		},
// 	)

// 	if err != nil {
// 		return userspb.CheckAuthResponse{
// 			Authenticated: false,
// 		}, err
// 	}

// 	return  userspb.CheckAuthResponse{
// 			Authenticated: res,
// 	}, nil
// }

func (s *Server) GetUserByLogin(
	ctx context.Context,
	req *userspb.GetUserByLoginRequest,
) (*userspb.GetUserByLoginResponse, error) {

	user, err := s.config.UcGetUserByLogin.GetUserByLogin(
		ctx,
		req.GetLogin(),
	)
	if err != nil {
		return nil, err
	}

	return &userspb.GetUserByLoginResponse{
		UserId:       user.Id,
		Login:        user.Login,
		PasswordHash: user.PasswordHash,
	}, nil
}


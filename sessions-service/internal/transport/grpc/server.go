package grpc

import (
	"context"

	sessionspb "AuthService/contracts/sessions"
	"AuthService/sessions-service/internal/usecase"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ServerConfig struct {
	UcCreateSession *usecase.CreateSessionUC
	UcDeleteSession *usecase.DeleteSessionUC
	UcGetSession    *usecase.GetSessionUC
}

type Server struct {
	sessionspb.UnimplementedSessionsServiceServer

	config ServerConfig
}

func NewServer(config ServerConfig) *Server {
	return &Server{
		config: config,
	}
}

func (s *Server) CreateSession(
	ctx context.Context,
	req *sessionspb.CreateSessionRequest,
) (*sessionspb.CreateSessionResponse, error) {

	session, err := s.config.UcCreateSession.CreateSession(
		ctx,
		req.GetUserId(),
	)
	if err != nil {
		return nil, err
	}

	return &sessionspb.CreateSessionResponse{
		Session: &sessionspb.Session{
			SessionId: session.SessionId,
			UserId:    session.UserId,
			ExpiresAt: timestamppb.New(session.ExpiresAt),
		},
	}, nil
}

func (s *Server) GetSession(
	ctx context.Context,
	req *sessionspb.GetSessionRequest,
) (*sessionspb.GetSessionResponse, error) {

	session, err := s.config.UcGetSession.GetSession(
		ctx,
		req.GetSessionId(),
	)
	if err != nil {
		return nil, err
	}

	return &sessionspb.GetSessionResponse{
		Session: &sessionspb.Session{
			SessionId: session.SessionId,
			UserId:    session.UserId,
			ExpiresAt: timestamppb.New(session.ExpiresAt),
		},
	}, nil
}

func (s *Server) DeleteSession(
	ctx context.Context,
	req *sessionspb.DeleteSessionRequest,
) (*emptypb.Empty, error) {

	err := s.config.UcDeleteSession.DeleteSession(
		ctx,
		req.GetSessionId(),
	)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
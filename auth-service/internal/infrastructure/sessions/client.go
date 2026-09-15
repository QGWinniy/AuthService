package clientsessions

import (
	"context"
	"AuthService/auth-service/internal/domain"
	sessionspb "AuthService/contracts/sessions"

	"google.golang.org/grpc"
)

type ClientSessionsService struct {
	client sessionspb.SessionsServiceClient
}

func NewClient(conn grpc.ClientConnInterface) *ClientSessionsService {
	return &ClientSessionsService{
		client: sessionspb.NewSessionsServiceClient(conn),
	}
}

func (c *ClientSessionsService) CreateSession(
	ctx context.Context,
	userID uint64,
) (*domain.Session, error) {

	response, err := c.client.CreateSession(
		ctx,
		&sessionspb.CreateSessionRequest{
			UserId: userID,
		},
	)
	if err != nil {
		return nil, err
	}

	return &domain.Session{
		SessionID: response.GetSession().GetSessionId(),
		UserID:    response.GetSession().GetUserId(),
		ExpiresAt: response.GetSession().GetExpiresAt().AsTime(),
	}, nil
}

func (c *ClientSessionsService) GetSession(
	ctx context.Context,
	sessionID string,
) (*domain.Session, error) {

	response, err := c.client.GetSession(
		ctx,
		&sessionspb.GetSessionRequest{
			SessionId: sessionID,
		},
	)
	if err != nil {
		return nil, err
	}

	return &domain.Session{
		SessionID: response.GetSession().GetSessionId(),
		UserID:    response.GetSession().GetUserId(),
		ExpiresAt: response.GetSession().GetExpiresAt().AsTime(),
	}, nil
}



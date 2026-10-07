package serverserviceclient

import (
	"AuthService/ssh-connect-service/internal/domain"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type ServerServiceClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewServerServiceClient(baseURL string, httpClient *http.Client) *ServerServiceClient {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: 5 * time.Second,
		}
	}

	return &ServerServiceClient{
		baseURL:    baseURL,
		httpClient: httpClient,
	}
}

func (c *ServerServiceClient)GetUserServerByID(
	ctx context.Context,
	serverID int,
	JWT string, // полная архитектурная, хуйня по факту мы уже проверили, что юзер авторизован, но микросервис userServer ждёт JWT, править нужно именно userServer и потом не прокидывать эту хуйню 
	// ну зато дохуя безопасно :)
) (*domain.UserServer, error) {

	getUserServerByIDRequest := getUserServerByIDRequest{
		UserServerID: serverID,
	}

	body, err := json.Marshal(getUserServerByIDRequest)

	u, err := url.JoinPath(
		c.baseURL,
		"/server",
	)

	if err != nil {
        return nil, fmt.Errorf("build URL: %w", err)
    }

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		u,
		bytes.NewReader(body),
	)

	 if err != nil {
        return nil, fmt.Errorf("create request: %v", err)
    }

    req.AddCookie(&http.Cookie{
        Name:  domain.CookieSessionJWT,
        Value: JWT,
    })

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return nil, fmt.Errorf("request server service: %v", err)
    }

    defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error StatusCode %v", resp)
	}

	var dto userServerResponse

	if err := json.NewDecoder(resp.Body).Decode(&dto); err != nil {
		return nil, fmt.Errorf("decode get user server response: %v", err)
	}

	return dto.toDomain(), nil
}
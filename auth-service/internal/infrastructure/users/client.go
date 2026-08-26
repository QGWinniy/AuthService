package clientusers

import (
	"context"
	"reg/auth-service/internal/domain"
	userspb "reg/contracts/users"

	"google.golang.org/grpc"
)

type ClientUsersService struct {
    client userspb.UserServiceClient
}

func NewClient(conn grpc.ClientConnInterface) *ClientUsersService {
    return &ClientUsersService{
        client: userspb.NewUserServiceClient(conn),
    }
}

func (c *ClientUsersService) CreateUser(
    ctx context.Context,
    login string,
    passwordHash string,
) (uint64, error) {

    response, err := c.client.CreateUser(
        ctx,
        &userspb.CreateUserRequest{
            Login:        login,
            PasswordHash: passwordHash,
        },
    )

    if err != nil {
        return 0, err
    }

    return response.GetUserId(), nil
}

// func (c *ClientUserService) GetUser(
//     ctx context.Context,
//     userId uint64,
// ) (*domain.User, error) {
//     response, err := c.client.GetUser(
//         ctx,
//         &userspb.GetUserRequest{
//             UserId: userId,
//         },
//     )

//     if err != nil {
//         return nil, err
//     }

//     user := &domain.User{
//         Id: userId,
//         Login: response.GetUser().Login,
//     }

//     return user, nil
// }

func (c *ClientUsersService) DeleteUser(
    ctx context.Context,
    userId uint64,
) error {
    _, err := c.client.DeleteUser(
        ctx,
        &userspb.DeleteUserRequest{
            UserId: userId,
        },
    )

    if err != nil {
        return err
    }

    return nil;
}

func (c *ClientUsersService) GetUserByLogin(
	ctx context.Context,
	login string,
) (*domain.User, error) {

	response, err := c.client.GetUserByLogin(
		ctx,
		&userspb.GetUserByLoginRequest{
			Login: login,
		},
	)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		Id:           response.GetUserId(),
		Login:        response.GetLogin(),
		PasswordHash: response.GetPasswordHash(),
	}, nil
}


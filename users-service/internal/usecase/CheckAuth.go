package usecase

// import (
// 	"context"
// 	"reg/users-service/internal/domain"
// 	"reg/users-service/internal/repository"
// )

// type CheckAuthConfig struct {
// 	UsersRepo repository.UsersRepository
// }

// type CheckAuthUC struct {
// 	config CheckAuthConfig
// }

// func NewCheckAuthUC(conf CheckAuthConfig) *CheckAuthUC {
// 	return &CheckAuthUC{
// 		config: conf,
// 	}
// }

// func (u *CheckAuthUC) CheckAuth(
// 	ctx context.Context,
// 	user *domain.User,
// ) (bool, error) {
// 	res, err := u.config.UsersRepo.CheckAuth(ctx, user);

// 	if err != nil {
// 		return false, err
// 	}

// 	return res, nil;
// }
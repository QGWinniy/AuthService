package repository

import "ssh/internal/domain"

type UserServerRepository interface {
	GetById(idServer int) (*domain.UserServer, error)
}
package repository

// import (
// 	"ssh/internal/domain"

// 	"golang.org/x/crypto/ssh"
// )

// type UserServerSessionRepository interface { // это именно конект к серверу и ничто иное 
// 	GetSession(userServer *domain.UserServer) (*ssh.Session, error) // в функции проверяем есть ли сессия если нет то создаём
// 	DeleteSession(userServer *domain.UserServer) 
// }
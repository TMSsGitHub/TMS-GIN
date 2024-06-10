package service

import "sync"

type UserService struct {
}

var (
	userService *UserService
	userOnce    sync.Once
)

func GetUserService() *UserService {
	userOnce.Do(func() {
		userService = &UserService{}
	})
	return userService
}

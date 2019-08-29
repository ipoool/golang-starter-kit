package interfaces

import "github.com/ipoool/golang-starter-kit/models"

// IUserRepository - UserRepository interface
type IUserRepository interface {
	GetUserByID(userID int) (user models.UserModel, err error)
}

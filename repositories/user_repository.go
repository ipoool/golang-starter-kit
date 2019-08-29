package repositories

import (
	"github.com/ipoool/golang-starter-kit/interfaces"
	"github.com/ipoool/golang-starter-kit/models"
)

// UserRepository - User repository struct
type UserRepository struct {
	DB interfaces.ISQL
}

// GetUserByID - For getting user by ID
func (r *UserRepository) GetUserByID(userID int) (user models.UserModel, err error) {
	db, err := r.DB.GetConnect()
	if err != nil {
		return
	}

	row, err := db.Prepare("SELECT id, name, email, is_active FROM users WHERE id = ?")
	if err != nil {
		return
	}
	err = row.QueryRow(userID).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.IsActive,
	)
	if err != nil {
		return
	}
	return
}

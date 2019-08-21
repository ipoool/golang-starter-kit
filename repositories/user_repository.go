package repositories

import (
	"github.com/ipoool/laporcuranmor-api/interfaces"
	"github.com/ipoool/laporcuranmor-api/models"
)

// UserRepository - User repository
type UserRepository struct {
	DB interfaces.ISQL
}

// GetUser - Get user by ID
func (r *UserRepository) GetUser(userID int) (*models.UserModel, error) {
	var user models.UserModel
	var err error

	db, err := r.DB.GetConnect()
	if err != nil {
		return nil, err
	}

	row, err := db.Prepare("SELECT * FROM users WHERE id = ?")
	if err != nil {
		return nil, err
	}
	err = row.QueryRow(userID).Scan(&user)
	if err != nil {
		return nil, err
	}

	return &user, err
}

// GetUsers - Get all Users using filter
func (r *UserRepository) GetUsers(filters map[string]string) ([]models.UserModel, error) {
	var users []models.UserModel
	var err error
	var query string

	db, err := r.DB.GetConnect()
	if err != nil {
		return nil, err
	}

	query = "SELECT * FROM users ORDER BY ? ?"

	rows, err := db.Query(query, filters["order"], filters["sort"])
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user models.UserModel
		err = rows.Scan(&user)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, err
}

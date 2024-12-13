package data

import (
	"errors"
	"online-shop-API/internal/types"
)

func (repo *Repository) Authorize(username, password string) (types.User, error) {
	var user types.User
	if err := repo.db.Model(&types.User{}).
		Preload("Role").
		Where("login = ? AND pass_hash = ?", username, password).
		First(&user).Error; err != nil {
		return types.User{}, errors.New("user not found")
	}
	return user, nil
}

func (repo *Repository) RegistrationUser(username string, password string, email string, role types.Role) error {
	var existingUser types.User
	if err := repo.db.Where("login = ?", username).First(&existingUser).Error; err == nil {
		return errors.New("user already exists")
	}

	var dbRole types.Role
	if err := repo.db.Where("role_name = ?", role.RoleName).First(&dbRole).Error; err != nil {
		return errors.New("role not found")
	}

	newUser := types.User{
		Login:    username,
		PassHash: password,
		Email:    email,
	}

	if err := repo.db.Create(&newUser).Error; err != nil {
		return errors.New("failed to create user")
	}

	if err := repo.db.Exec("INSERT INTO user_role VALUES (?, ?)", newUser.UserId, dbRole.RoleID).Error; err != nil {
		return errors.New("failed to create user")
	}

	return nil
}

func (repo *Repository) GetRole(roleName string) (*types.Role, error) {
	var role types.Role
	if err := repo.db.Where("role_name = ?", roleName).First(&role).Error; err != nil {
		return nil, errors.New("role not found")
	}

	return &role, nil
}

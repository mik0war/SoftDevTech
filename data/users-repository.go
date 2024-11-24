package data

import (
	"errors"
	"online-shop-API/types"
)

var users = map[string]types.User{
	"admin": {
		Username: "admin",
		Password: "root",
		Roles: []types.Role{
			{Name: "admin"},
			{Name: "user"},
		},
	},
}

func Authorize(username, password string) (types.User, error) {
	if users[username].Equal(types.User{Username: username, Password: password}) {
		return users[username], nil
	} else {
		return types.User{}, errors.New("user not found")
	}
}

func RegistrationUser(username string, password string, role types.Role) error {
	if _, exists := users[username]; exists {
		return errors.New("user already exists")
	}
	users[username] = types.User{Username: username, Password: password, Roles: []types.Role{role}}
	return nil
}

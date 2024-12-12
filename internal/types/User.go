package types

import "github.com/dgrijalva/jwt-go"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Roles    []Role `json:"role"`
}

func (u User) Equal(other User) bool {
	return u.Username == other.Username && u.Password == other.Password
}

type Role struct {
	Name string `json:"name"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Roles    []Role `json:"role"`
}

type Claims struct {
	Username string `json:"username"`
	Roles    []Role `json:"role"`
	jwt.StandardClaims
}

type JwtResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

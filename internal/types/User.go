package types

import "github.com/dgrijalva/jwt-go"

type User struct {
	UserId   int64  `gorm:"primaryKey;autoIncrement"`
	Login    string `json:"username"`
	PassHash string `json:"password"`
	Email    string `json:"email"`

	Role []Role `gorm:"many2many:user_role"`
}

type Role struct {
	RoleName    string `json:"name"`
	RoleID      uint   `gorm:"primaryKey" json:"id"`
	AccessLevel uint   `json:"access_level"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
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

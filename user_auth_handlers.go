package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"online-shop-API/internal/types"
	"slices"
	"time"
)

var jwtKey = []byte("my_secret_key")

func generateToken(username string, roles []types.Role, expirationTime time.Time) (string, error) {
	claims := &types.Claims{
		Username: username,
		Roles:    roles,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// @Summary Регистрация нового пользователя
// @Description Регистрирует нового пользователя в системе
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body types.Credentials true "Учётные данные пользователя"
// @Success 201 {object} types.SuccessResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /auth/register [post]
func (handler *Handler) registration(c *gin.Context) {
	var credentials types.Credentials
	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "invalid request"})
		return
	}

	err := handler.productRepo.RegistrationUser(credentials.Username, credentials.Password, credentials.Email, types.Role{RoleName: "User"})

	if err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, types.SuccessResponse{Message: "User created"})
}

// @Summary Вход пользователя
// @Description Авторизует пользователя и возвращает токены
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body types.Credentials true "Учётные данные пользователя"
// @Success 200 {object} types.JwtResponse
// @Failure 400 {object} types.ErrorResponse
// @Failure 401 {object} types.ErrorResponse
// @Router /auth/login [post]
func (handler *Handler) login(c *gin.Context) {
	var credentials types.Credentials
	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "invalid request"})
		return
	}

	var user types.User
	var err error
	if user, err = handler.productRepo.Authorize(credentials.Username, credentials.Password); err != nil {
		c.JSON(http.StatusUnauthorized, types.ErrorResponse{Error: "unauthorized"})
		return
	}

	credentials.Roles = user.Role

	accessToken, err := generateToken(
		credentials.Username,
		credentials.Roles,
		time.Now().Add(1*time.Minute),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "could not create token"})
		return
	}

	refreshToken, err := generateToken(
		credentials.Username,
		credentials.Roles,
		time.Now().Add(15*time.Hour),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "could not create token"})
		return
	}

	c.JSON(http.StatusOK, types.JwtResponse{AccessToken: accessToken, RefreshToken: refreshToken})
}

// @Summary Обновить токен доступа
// @Description Обновляет токен доступа по действующему refresh-токену
// @Tags auth
// @Produce json
// @Param Authorization header string true "Токен для обновления"
// @Success 200 {object} types.SuccessResponse
// @Failure 401 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /auth/refresh [post]
func refresh(c *gin.Context) {

	tokenString := c.Request.Header.Get("Authorization")
	claims := &types.Claims{}
	refreshToken, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if claims.ExpiresAt < time.Now().Unix() {
		c.JSON(http.StatusUnauthorized, types.ErrorResponse{Error: "token expired"})
		c.Abort()
		return
	}

	if err != nil || !refreshToken.Valid {
		c.JSON(http.StatusUnauthorized, types.ErrorResponse{Error: "unauthorized"})
		c.Abort()
		return
	}

	token, err := generateToken(claims.Username, claims.Roles, time.Now().Add(1*time.Minute))

	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "could not create token"})
		return
	}

	c.JSON(http.StatusOK, types.SuccessResponse{Message: token})
}

func authMiddleware(requiredRole types.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		claims := &types.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if claims.ExpiresAt < time.Now().Unix() {
			c.JSON(http.StatusUnauthorized, types.ErrorResponse{Error: "token expired"})
			c.Abort()
			return
		}

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, types.ErrorResponse{Error: "unauthorized"})
			c.Abort()
			return
		}

		if !slices.Contains(claims.Roles, requiredRole) {
			c.JSON(http.StatusForbidden,
				types.ErrorResponse{Error: "forbidden: insufficient permissions"})
			c.Abort()
			return
		}
		c.Next()
	}
}

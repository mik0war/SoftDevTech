package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"online-shop-API/data"
	"online-shop-API/types"
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

// @Summary Register a new user
// @Description Create a new user with "user" role
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param request body types.UserAuthData true "User credentials"
// @Success 201 {object} types.SuccessResponse "User created successfully"
// @Failure 400 {object} types.ErrorResponse "Invalid request or user already exists"
// @Router /register [post]
func registration(c *gin.Context) {
	var credentials types.Credentials
	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "invalid request"})
		return
	}

	err := data.RegistrationUser(credentials.Username, credentials.Password, types.Role{Name: "user"})

	if err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, types.SuccessResponse{Message: "User created"})
}

// @Summary Login user
// @Description Authenticate a user and return access and refresh tokens
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param request body types.UserAuthData true "User credentials"
// @Success 200 {object} types.JwtResponse "Tokens generated successfully"
// @Failure 400 {object} types.ErrorResponse "Invalid request"
// @Failure 401 {object} types.ErrorResponse "Unauthorized"
// @Failure 500 {object} types.ErrorResponse "Internal server error"
// @Router /login [post]
func login(c *gin.Context) {
	var credentials types.Credentials
	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "invalid request"})
		return
	}

	var user types.User
	var err error
	if user, err = data.Authorize(credentials.Username, credentials.Password); err != nil {
		c.JSON(http.StatusUnauthorized, types.ErrorResponse{Error: "unauthorized"})
		return
	}

	credentials.Roles = user.Roles

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
		time.Now().Add(15*time.Minute),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "could not create token"})
		return
	}

	c.JSON(http.StatusOK, types.JwtResponse{AccessToken: accessToken, RefreshToken: refreshToken})
}

// @Summary Refresh token
// @Description Refresh the access token using a valid refresh token
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Refresh_token"
// @Success 200 {object} types.SuccessResponse "Token refreshed successfully"
// @Failure 401 {object} types.ErrorResponse "Token expired or unauthorized"
// @Failure 500 {object} types.ErrorResponse "Internal server error"
// @Router /refresh [post]
func refresh(c *gin.Context) {

	tokenString := c.Request.Header.Get("Authorization")
	claims := &types.Claims{}
	refreshToken, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if refreshToken == nil {
		c.JSON(http.StatusUnauthorized, types.ErrorResponse{Error: "empty token"})
		c.Abort()
		return
	}

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

		if token == nil {
			c.JSON(http.StatusUnauthorized, types.ErrorResponse{Error: "empty token"})
			c.Abort()
			return
		}

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

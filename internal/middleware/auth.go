package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var (
	SecretKey       = []byte("s3cr3t")
	accessTokenExp  = 7 * 24 * time.Minute
	refreshTokenExp = 7 * 24 * time.Hour
)

func GenerateAccessToken() (string, error) {
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(accessTokenExp).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SecretKey)
}

func GenerateRefreshToken() (string, error) {
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(refreshTokenExp).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SecretKey)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return SecretKey, nil
		})
		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
	}
}

func RefreshTokenHandler(c *gin.Context) {
	// extract refresh token from request
	refreshToken := c.PostForm("refresh_token")
	if refreshToken == "" {
		c.JSON(400, gin.H{"error": "refresh token required"})
		c.Abort()
		return
	}
	// validate refresh token
	token, err := jwt.ParseWithClaims(refreshToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil || !token.Valid {
		c.JSON(401, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}
	// generate new access token if refresh token is valid
	accessToken, err := GenerateAccessToken()
	if err != nil {
		c.JSON(500, gin.H{"error": "internal server error"})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"access_token": accessToken})
}

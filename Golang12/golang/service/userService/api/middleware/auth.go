package middleware

import (
	"errors"
	"net/http"
	"source-base-go/golang/service/userService/config"
	"source-base-go/golang/service/userService/infrastructure/repository/util"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(verifier util.Verifier) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
			c.Abort()
			return
		}
		var jwtSecretKey = []byte(config.GetString("jwt.secretKey"))

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid signing method")
			}
			return jwtSecretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token err"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userId := int(claims["user_id"].(float64))

			tokenExists, err := verifier.Verifier(userId, tokenString)
			if err != nil || tokenExists {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or revoked token"})
				c.Abort()
				return
			}
			c.Set("user_id", userId)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid claims"})
			c.Abort()
			return
		}

		c.Next()
	}
}

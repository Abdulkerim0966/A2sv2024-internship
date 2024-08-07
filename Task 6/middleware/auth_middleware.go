package middleware

import (
	"fmt"
	
	"strings"
	"task/controllers"
	"task/models"
	

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMidddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(401, gin.H{"error": "Authorization header is required"})
            c.Abort()
            return
        }

        authParts := strings.Split(authHeader, " ")
        if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
            c.JSON(401, gin.H{"error": "Invalid authorization header"})
            c.Abort()
            return
        }

        tokenString := authParts[1]
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
            }
            return controllers.JwtSecret, nil
        })

        if err != nil || !token.Valid {
            c.JSON(401, gin.H{"error": "Invalid JWT"})
            c.Abort()
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            c.JSON(401, gin.H{"error": "Invalid JWT claims"})
            c.Abort()
            return
        }

        // Set user data in context
        user := models.User{
            UserName: claims["userName"].(string),
            Role:     claims["role"].(bool),
        }
        c.Set("user", user)
        c.Next()
    }
}
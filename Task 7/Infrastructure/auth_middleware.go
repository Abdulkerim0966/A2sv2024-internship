package Infrastructure

import (
	"net/http"
	"task/Domain"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMidddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if strings.ToLower(authParts[0]) != "bearer" {
			c.JSON(401, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}

		if len(authParts) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		tokenString := authParts[1]
		authorized, err := VerifyAccessToken(tokenString, secret)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Error verifying token"})
			c.Abort()
			return
		}
		if !authorized {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		claims, err := ExtractClaims(tokenString, secret)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Error extracting claims"})
            c.Abort()
            return
        }      
		clm := Domain.JwrCustonClaims{
			UserName: claims["username"].(string),
			Role:     claims["role"].(bool),
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: int64(claims["exp"].(float64)),
			}}

		c.Set("claims", clm)

		c.Next()

	}
}


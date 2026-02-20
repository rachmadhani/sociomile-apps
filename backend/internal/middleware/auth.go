package middleware

import (
	"net/http"
	"sociomile-apps/config"
	"sociomile-apps/internal/database"
	authDTO "sociomile-apps/internal/dto/auth"
	model "sociomile-apps/internal/models"
	"sociomile-apps/internal/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user *model.User) (string, error) {
	expirationTime := time.Now().Add(time.Hour * time.Duration(config.AppConfig.JWTExpirationHours))

	claims := authDTO.Claims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   string(user.Role),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   user.ID.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.AppConfig.JWTSecret))
}

func ValidateToken(tokenString string) (*authDTO.Claims, error) {
	claims := &authDTO.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.JWTSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.Unauthorized(c, "Authorization header required")
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.Unauthorized(c, "Invalid authorization header format")
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims, err := ValidateToken(tokenString)
		if err != nil {
			utils.Unauthorized(c, "Invalid or expired token")
			c.Abort()
			return
		}

		var user model.User
		if err := database.DB.First(&user, "id = ?", claims.UserID).Error; err != nil {
			utils.Unauthorized(c, "User not found")
			c.Abort()
			return
		}

		c.Set("user", &user)
		c.Set("user_id", user.ID)

		c.Next()
	}
}

func RequireRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Role not found in context"})
			c.Abort()
			return
		}

		roleStr := userRole.(string)
		isAuthorized := false
		for _, r := range roles {
			if r == roleStr {
				isAuthorized = true
				break
			}
		}

		if !isAuthorized {
			c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to access this resource"})
			c.Abort()
			return
		}

		c.Next()
	}
}

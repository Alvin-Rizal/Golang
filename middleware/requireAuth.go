// RequireAuth is a middleware function for authentication.
package middleware

import (
	"fmt"
	"golang-mnc/initializer"
	"golang-mnc/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		fmt.Println("Error retrieving Authorization cookie:", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	fmt.Println("Received token:", tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || method != jwt.SigningMethodHS256 {

			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil || !token.Valid {
		fmt.Println("Error validating token:", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Error parsing token claims")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	expirationTime := claims["exp"].(float64)
	currentTime := float64(time.Now().Unix())
	if currentTime > expirationTime {
		fmt.Println("Token expired")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		fmt.Println("Sub claim not found in the token")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var user models.User
	if err := initializer.DB.First(&user, sub).Error; err != nil {
		fmt.Println("Error querying user from the database:", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if user.ID == 0 {
		fmt.Println("User not found in the database")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	fmt.Println("Authenticated user:", user)
	c.Set("user", user)
	c.Next()
}

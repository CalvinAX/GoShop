package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/CalvinAX/GoShop/initializers"
	"github.com/CalvinAX/GoShop/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	tokenString := c.GetHeader("X-AUTH-TOKEN")

	if tokenString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "No token sent!",
		})

		return
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("WARNING ##### Wrong singning method!")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("AUTH_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token expired!",
			})
		}

		var user models.User
		initializers.DB.First(&user, "id = ?", claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "User does not exist!",
			})
		}

		c.Set("user", user)

		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Not a valid token!",
		})
	}
}

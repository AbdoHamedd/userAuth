package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Auth(c *gin.Context) {
	// todo get the Token
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	count := 0
	for _, i2 := range tokenString {
		if i2 == '.' {
			count++
		}
	}
	if count != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	//todo Decode & Validate it
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("1a2b3d4o5h6a7m8e6d"), nil
	})

	// todo Check The Exp
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > float64(claims["exp"].(float64)) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		//todo Continue
		c.Next()
	}
	c.AbortWithStatus(http.StatusUnauthorized)

}

package middleware

import (
	"errors"
	"go/ems/domain"
	"go/ems/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func extractBearerToken(jwt string) string {
	if len(strings.Split(jwt, " ")) == 2 {
		return strings.Split(jwt, " ")[1]
	}
	return ""
}

func AuthUsersCheck(c *gin.Context) {
	// Extract token from Header
	tokenString := extractBearerToken(c.GetHeader("Authorization"))
	if tokenString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, domain.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   "Login Required",
		})
		return
	}

	// Verify token validity
	token, err := verifyJWTToken(tokenString)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		})
		return
	}

	_, OK := token.Claims.(jwt.MapClaims)
	if !OK {
		c.AbortWithStatusJSON(http.StatusInternalServerError, domain.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   "Unable to parse claim",
		})
		return
	}

	c.Next()
}

func verifyJWTToken(tokenString string) (*jwt.Token, error) {

	jwtKey := helper.LoadEnvFile("JWT_SECRECT_KEY")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, OK := token.Method.(*jwt.SigningMethodHMAC); !OK {
			return nil, errors.New("bad signed method received")
		}
		return []byte(jwtKey), nil
	})

	// if token.Valid {
	// 	return true, nil
	// } else if errors.Is(err, jwt.ErrTokenMalformed) {
	// 	return false, jwt.ErrTokenMalformed
	// } else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
	// 	return false, jwt.ErrTokenExpired
	// } else {
	// 	return false, jwt.ErrTokenNotValidYet
	// }

	return token, err
}

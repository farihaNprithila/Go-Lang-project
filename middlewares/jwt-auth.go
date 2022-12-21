package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gitlab.com/pragmaticreviews/golang-gin-poc/helper"
	"gitlab.com/pragmaticreviews/golang-gin-poc/service"
	"log"
	"net/http"
)

/*Validate the given token*/
func AuthorizeJWT(jwtservice service.JWTService) gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.CreateCommonErrorResponse(
				"Failed to process request", "No token found", nil)
			context.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		token, err := jwtservice.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]:", claims["user_id"])
			log.Println("Claim[issuer]:", claims["issuer"])
		} else {
			log.Println(err)
			response := helper.CreateCommonErrorResponse(
				"Token is not valid", err.Error(), nil)
			context.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}

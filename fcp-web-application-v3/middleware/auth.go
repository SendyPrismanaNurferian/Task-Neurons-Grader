package middleware

import (
	"net/http"
	"a21hc3NpZ25tZW50/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		sessionToken, err := ctx.Request.Cookie("session_token")
		if err != nil {
			if ctx.GetHeader("Content-Type") == "application/json" {
				ctx.AbortWithStatus(http.StatusUnauthorized)
			} else {
				ctx.Redirect(http.StatusSeeOther, "/login")
			}
			return
		}

		tkn := sessionToken.Value
		claims := &model.Claims{}
		token, err := jwt.ParseWithClaims(tkn, claims, func(token *jwt.Token) (interface{}, error) {
			return model.JwtKey, nil
		})
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if !token.Valid {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("email", claims.Email)

		ctx.Next()
	})
}

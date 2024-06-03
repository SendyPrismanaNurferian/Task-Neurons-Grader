package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		// Periksa keberadaan cookie session_token
		cookie, err := ctx.Request.Cookie("session_token")
		if err != nil {
			if ctx.GetHeader("Content-Type") == "application/json" {
				ctx.AbortWithStatus(http.StatusUnauthorized)
			} else {
				ctx.Redirect(http.StatusSeeOther, "/login")
			}
			return
		}

		tokenString := cookie.Value
		claims := &model.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
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

		ctx.Set("id", claims.UserID)

		ctx.Next() // TODO: answer here
	})
}

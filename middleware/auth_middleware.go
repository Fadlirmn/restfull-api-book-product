package middleware

import(
	"go-roadmap/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)


func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader:= c.GetHeader("Authorization")

		if authHeader=="" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":"Authorization header required",
			})
			c.Abort()
			return 
		}
		tokenString := strings.Split(authHeader,"")
		token, err:= jwt.ParseWithClaims(tokenString[1],&utils.Claims{},func(t *jwt.Token) (interface{}, error) {
			return utils.SECRET_KEY,nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized,gin.H{
				"error":"invalid token ",
			})
			c.Abort()
			return 
		}

		claims:= token.Claims.(*utils.Claims)

		c.Set("id",claims.UserID)
		c.Set("role",claims.Role)

		c.Next()
	} 
	
	
}

package main

import (
	"kanbanmusume_ura/db"
	"kanbanmusume_ura/models"
	"kanbanmusume_ura/services"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var User models.User

func main() {
	db.Init()
	router := makeRouter()

	router.Run(":3000")
}

func makeRouter() *gin.Engine {
	router := gin.Default()
	authMiddleware := makeAuthMiddleware()
	router.POST("/login", authMiddleware.LoginHandler)

	return router
}

var ikey = "name"

func makeAuthMiddleware() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "myRealm",
		Key:             []byte("secretKey"),
		Timeout:         24 * time.Hour,
		MaxRefresh:      24 * time.Hour,
		IdentityKey:     ikey,
		PayloadFunc:     makePayload,
		IdentityHandler: identityHandler,
		Authenticator:   authenticator,
		Unauthorized:    makeLoginFailedMassage,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	})

	if err != nil {
		return nil
	}
	errInit := authMiddleware.MiddlewareInit()
	if errInit != nil {
		return nil
	}

	return authMiddleware
}

func makePayload(data interface{}) jwt.MapClaims {
	if v, ok := data.(*models.User); ok {
		return jwt.MapClaims{
			ikey: v.Name,
		}
	}
	return jwt.MapClaims{}
}

func identityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &models.User{
		Name: claims[ikey].(string),
	}
}

func authenticator(c *gin.Context) (interface{}, error) {
	var loginInfo models.LoginRequest
	if err := c.ShouldBind(&loginInfo); err != nil {
		panic(err)
		return "", jwt.ErrMissingLoginValues
	}
	name := loginInfo.Username
	password := loginInfo.Password
	service := services.UserService{}
	user, err := service.FindByName(name)
	if err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	// if err = service.IsSamePassword(user.Token, password); err != nil {
	// 	return nil, jwt.ErrFailedAuthentication
	// }
	if password != "password" {
		return nil, jwt.ErrFailedAuthentication
	}

	return user, nil
}

func makeLoginFailedMassage(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

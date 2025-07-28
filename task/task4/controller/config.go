package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	OK         int = 200
	PARAM_ERR  int = 400
	AUTH_ERR   int = 401
	SERVER_ERR int = 500
)

const (
	TOKEN          string = "Authorization"
	JWT_SECRET_KEY string = "your_secret_key"
	USER_ID        string = "userId"
)

var Router *gin.Engine

type Response struct {
	Code int
	Msg  string
	Data interface{}
}

func init() {
	// gin
	Router = gin.Default()
	Router.Any("/test", func(c *gin.Context) {
		Ok(c, nil)
	})
	UserApi()
	PostApi()
	CommentApi()
}

func auth(c *gin.Context) {
	// 获取token
	tokenString := c.GetHeader(TOKEN)
	if tokenString == "" {
		Err(c, AUTH_ERR, "用户未登录")
		c.Abort()
		return
	}
	// 解析token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET_KEY), nil
	})
	if err != nil {
		Err(c, SERVER_ERR, "验签失败")
		c.Abort()
		return
	}
	// 验证token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := uint(claims["id"].(float64))
		exp := int64(claims["exp"].(float64))
		if exp < time.Now().Unix() {
			Err(c, AUTH_ERR, "登录信息已失效")
			c.Abort()
			return
		}
		c.Set("userId", userId)
	}
	// 放行
	c.Next()
}

func Ok(c *gin.Context, data interface{}) {
	c.JSON(OK, Response{
		Code: OK,
		Msg:  "ok",
		Data: data,
	})
}

func Err(c *gin.Context, code int, msg string) {
	c.JSON(code, Response{
		Code: code,
		Msg:  msg,
	})
}

func getUserId(c *gin.Context) uint {
	userId, ok := c.Get(USER_ID)
	if ok {
		return userId.(uint)
	}
	return 0
}

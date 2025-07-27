package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"task4/req"
	"task4/service"
	"time"
)

func Register(c *gin.Context) {
	var userAdd req.UserAdd
	if err := c.ShouldBind(&userAdd); err != nil {
		Err(c, PARAM_ERR, "参数错误")
		return
	}
	// 用户名防重复校验
	user := service.UserFindByUsername(userAdd.Username)
	if user != nil {
		Err(c, PARAM_ERR, "用户已存在")
		return
	}
	// 邮箱防重复校验
	count2 := service.UserCountByEmail(userAdd.Email)
	if count2 > 0 {
		Err(c, PARAM_ERR, "用户已存在")
		return
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userAdd.Password), bcrypt.DefaultCost)
	if err != nil {
		Err(c, PARAM_ERR, "密码加密异常")
		return
	}
	userAdd.Password = string(hashedPassword)
	// 添加用户
	service.UserAdd(&userAdd)
	// 响应
	Ok(c, nil)
}

func Login(c *gin.Context) {
	var userLogin req.UserLogin
	if err := c.ShouldBind(&userLogin); err != nil {
		Err(c, PARAM_ERR, "参数错误")
		return
	}
	// 查询用户
	user := service.UserFindByUsername(userLogin.Username)
	if user == nil {
		Err(c, PARAM_ERR, "用户不存在")
		return
	}
	// 校验密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password)); err != nil {
		Err(c, PARAM_ERR, "密码错误")
		return
	}
	// 生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(JWT_SECRET_KEY))
	if err != nil {
		Err(c, SERVER_ERR, "生成token错误")
		return
	}
	// 响应
	Ok(c, tokenString)
}

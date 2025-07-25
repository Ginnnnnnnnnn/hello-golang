package aop

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Aop(c *gin.Context) {
	fmt.Println("aop before")
	c.Next()
	fmt.Println("aop after")
}

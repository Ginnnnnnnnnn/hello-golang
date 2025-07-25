package mod

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"slices"
)

type User struct {
	Id    int    `form:"id" json:"id" binding:"required,number"`
	Name  string `form:"name" json:"name" binding:"required"`
	Sex   int    `form:"sex" json:"sex" binding:"sex"`
	Email string `form:"email" json:"email" binding:"omitempty,email"`
}

var sex validator.Func = func(fl validator.FieldLevel) bool {
	sex := fl.Field().Int()
	if !slices.Contains([]int64{0, 1, 2}, sex) {
		return false
	}
	return true
}

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("sex", sex)
	}
}

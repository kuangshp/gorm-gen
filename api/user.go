package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm_gen/dao"
	"gorm_gen/model"
	"gorm_gen/utils"
)

func GetAllUser(ctx *gin.Context) {
	//first, _ := query.User.Find()
	find, _ := utils.DB.User.Find()
	ctx.JSON(200, find)
}

func DeleteUser(ctx *gin.Context) {
	info, err := dao.User.Where(dao.User.ID.Eq(1)).Delete()
	fmt.Println(info, err)
	ctx.JSON(200, info)
}

// CreateUser 创建用户
func CreateUser(ctx *gin.Context) {
	u := &model.User{
		Username: "谢逊",
		Age:      30,
		Password: "qwer1234",
		Phone:    "12345678901",
	}
	err := dao.User.Create(u)
	if err != nil {

	}
	ctx.JSON(200, err)
}

type IUser interface {
	GetAllUserApi(ctx *gin.Context)
}

type User struct {
	db *dao.Query
}

func (u User) GetAllUserApi(ctx *gin.Context) {
	result, _ := u.db.User.Find()
	ctx.JSON(200, result)
}

func NewUser(db *dao.Query) IUser {
	return &User{
		db: db,
	}
}

package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-bai/forward/database"
	"github.com/go-bai/forward/model"
	"github.com/go-bai/forward/pkg"
)

type RegisterReq struct {
	Username string `json:"username" binding:"required,min=1"`
	Password string `json:"password" binding:"required,min=1"`
}

func Register(ctx *gin.Context) {
	req := new(RegisterReq)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		ctx.Abort()
		return
	}

	if req.Username == "" || req.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "username and password cannot be empty",
		})
		ctx.Abort()
		return
	}

	user := &model.User{
		Username: req.Username,
		Password: req.Password,
	}

	if err := user.HashPassword(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		ctx.Abort()
		return
	}

	_, err := database.Engine.InsertOne(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"username": user.Username})
}

type LoginReq struct {
	Username string `json:"username" binding:"required,min=1"`
	Password string `json:"password" binding:"required,min=1"`
}

func Login(ctx *gin.Context) {
	req := new(LoginReq)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		ctx.Abort()
		return
	}

	user := new(model.User)
	has, err := database.Engine.Where("username = ?", req.Username).Get(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		ctx.Abort()
		return
	}

	if !has || user.CheckPassword(req.Password) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "incorrect username or password",
		})
		ctx.Abort()
		return
	}

	token, err := pkg.GenerateToken(user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"username": user.Username,
		"token":    token,
		"code":     http.StatusOK,
	})
}

type UserListResp struct {
	List []*model.User `json:"list"`
}

func UserList(ctx *gin.Context) {
	userList := make([]*model.User, 0)
	err := database.Engine.Find(&userList)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSONP(http.StatusOK, userList)
}

func UserDetail(ctx *gin.Context) {
	username := ctx.GetString("username")
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  fmt.Sprintf("hello, %s", username),
	})
}

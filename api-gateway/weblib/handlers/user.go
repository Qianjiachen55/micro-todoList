package handlers

import (
	"api-gateway/pkg/utils"
	"api-gateway/services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

// register
func UserRegister(ginCtx *gin.Context) {
	var userReq services.UserRequest

	PanicIfUserError(ginCtx.Bind(&userReq))

	userService := ginCtx.Keys["userService"].(services.UserService)

	userResp, err := userService.UserRegister(context.Background(), &userReq)

	PanicIfUserError(err)
	ginCtx.JSON(http.StatusOK, gin.H{"data": userResp})

}


// login
func UserLogin(ginCtx *gin.Context)  {
	var userReq services.UserRequest
	PanicIfUserError(ginCtx.Bind(&userReq))
	//从keys中取出service
	userService := ginCtx.Keys["userService"].(services.UserService)
	userResp ,err := userService.UserLogin(context.Background(),&userReq)
	PanicIfUserError(err)

	token ,err := utils.GenerateToken(uint(userResp.UserDetail.ID))

	ginCtx.JSON(http.StatusOK,gin.H{
		"code":userResp.Code,
		"msg": "success",
		"data": gin.H{
			"user":userResp.UserDetail,
			"token":token,
		},
	})
}
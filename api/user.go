package api

import (
	"fmt"
	"hextech/common/request"
	"hextech/common/response"
	"hextech/utils/jwt"

	"hextech/schema"

	"github.com/gin-gonic/gin"
)

// @Summary     用户注册
// @Description 新用户注册接口
// @Tags        用户接口
// @Accept      json
// @Produce     json
// @Param       request body schema.RegisterRequest true "注册信息"
// @Success     200 {object} map[string]interface{}
// @Failure     400 {object} map[string]interface{}
// @Router      /user/register [post]
func (ctrl *Controller) Register(c *gin.Context) {
	var req schema.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 这里应该添加用户注册逻辑
	response.OkWithMessage("注册成功", c)
}

// @Summary     用户登录
// @Description 用户登录接口
// @Tags        用户接口
// @Accept      json
// @Produce     json
// @Param       request body schema.RegisterRequest true "登录信息"
// @Success     200 {object} map[string]interface{}
// @Failure     400 {object} map[string]interface{}
// @Router      /user/login [post]
func (ctrl *Controller) Login(c *gin.Context) {
	var req schema.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 这里应该添加用户登录逻辑
	// 假设登录成功，创建token
	claims := ctrl.JWT.CreateClaims(jwt.BaseClaims{
		ID:       1,
		Username: req.Username,
	})
	token, err := ctrl.JWT.CreateToken(claims)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(map[string]any{"token": token}, "登录成功", c)
}

// @Summary     获取用户信息
// @Description 获取当前登录用户信息
// @Tags        用户接口
// @Produce     json
// @Success     200 {object} schema.UserInfo
// @Router      /user/info [get]
func (ctrl *Controller) GetUserInfo(c *gin.Context) {
	// 这里应该从token中获取用户信息
	username := request.GetUserName(c)

	userInfo := schema.UserInfo{
		Username: username,
	}

	response.OkWithDetailed(userInfo, "获取成功", c)
}

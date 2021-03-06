package controller

import (
	"github.com/gin-gonic/gin"
	"mkgo/common"
	"mkgo/middleware/jwtauth"
	"mkgo/model"
	"net/http"
)

type UserController struct{}

type loginResponse struct {
	Token string     `json:"token"`
	User  model.User `json:"user"`
}

var userManager = &model.UserManager{}

func (ctrl *UserController) Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	b := userManager.AddUser(&model.User{Name:username, Password: password})
	if !b {
		c.JSON(http.StatusOK, common.NewJSONResponse(common.CodeErrorRegister, "Register failed", nil))
		return
	}
	c.JSON(http.StatusOK, common.NewJSONResponse(common.CodeSuccess, "Register success", nil))

}

func (ctrl *UserController) Login(c *gin.Context) {
	token := jwtauth.GetToken(c, "mickyliu")
	if token == "" {
		result := common.NewJSONResponse(common.CodeErrorInternal, "create token failed", nil)
		c.JSON(common.CodeSuccess, result)
		return
	}
	loginResp := loginResponse{Token: token, User: model.User{Name: "mickyliu", Password: "123456"}}
	result := common.NewJSONResponse(common.CodeSuccess, common.MessageSuccess, loginResp)
	c.JSON(common.CodeSuccess, result)
}

func (ctrl *UserController) Logout(c *gin.Context) {
	c.String(common.CodeSuccess, "Logout!")
}

func (ctrl *UserController) GetUserById(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	if id != "" {
		user := userManager.GetUserById(id)
		result := common.NewJSONResponse(common.CodeSuccess, common.MessageSuccess, user)
		c.JSON(common.CodeSuccess, result)
	} else {
		c.JSON(common.CodeSuccess, common.NewEmptyDataResponse("user not find"))
	}
}

func (ctrl *UserController) GetUserList(c *gin.Context) {
	users := userManager.GetUserList()
	result := common.NewJSONResponse(common.CodeSuccess, common.MessageSuccess, users)
	c.JSON(common.CodeSuccess, result)
}

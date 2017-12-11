package controller

import (
	"github.com/gin-gonic/gin"
	"mkgo/common"
	"mkgo/middleware/jwtauth"
	"mkgo/model"
	"net/http"
	"mkgo/mkdb"
	"mkgo/mklog"
	"go.uber.org/zap"
)

type UserController struct{}

type loginResponse struct {
	Token string     `json:"token"`
	User  model.User `json:"user"`
}

func (ctrl UserController) Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	inertUser := `INSERT INTO user ( name, password) VALUES (?, ?)`
	_, err := mkdb.GetWriteDB().Exec(inertUser, username, password)
	if err != nil {
		mklog.Logger.Error("[register]", zap.Error(err))
		c.JSON(http.StatusOK, common.NewJSONResponse(common.CodeErrorRegister, "Register failed", nil))
		return
	}
	c.JSON(http.StatusOK, common.NewJSONResponse(common.CodeSuccess, "Register success", nil))

}

func (ctrl UserController) Login(c *gin.Context) {
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

func (ctrl UserController) Logout(c *gin.Context) {
	c.String(common.CodeSuccess, "Logout!")
}

func (ctrl UserController) GetUser(c *gin.Context) {
	id := c.Query("id")
	if id == "10001" {
		user := model.User{Name: "mickyliu", Password: "123456"}
		result := common.NewJSONResponse(common.CodeSuccess, common.MessageSuccess, user)
		c.JSON(common.CodeSuccess, result)
	} else {
		c.JSON(common.CodeSuccess, common.NewEmptyDataResponse("user not find"))
	}
}

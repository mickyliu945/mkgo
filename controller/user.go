package controller

import (
	"github.com/gin-gonic/gin"
	"mkgo/common"
	"mkgo/middleware/jwtauth"
	"mkgo/model"
	"net/http"
	"mkgo/mkdb"
)

type UserController struct{}

type loginResponse struct {
	Token string     `json:"token"`
	User  model.User `json:"user"`
}

func (ctrl UserController) Register(c *gin.Context) {
	inertUser := `INSERT INTO place (country, telcode) VALUES (?, ?)`
	mkdb.GetWriteDB().Exec(inertUser)
	c.String(http.StatusOK, "Register Success!")
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

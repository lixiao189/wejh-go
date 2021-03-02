package userController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/models"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/services/userCenterServices"
	"wejh-go/app/services/userServices"
	"wejh-go/app/utils"
	"wejh-go/app/utils/stateCode"
	"wejh-go/service/database"
)

type bindJHForm struct {
	UserName  string `json:"username"`
	PassWord  string `json:"password"`
	LoginType string `json:"type"`
}

func BindJHID(c *gin.Context) {
	var postForm bindJHForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		utils.JsonFailedResponse(c, stateCode.ParamError, nil)
		return
	}
	session, err := sessionServices.GetWechatSession(c)
	if err != nil {
		utils.JsonFailedResponse(c, stateCode.GetOpenIDFail, nil)
		return
	}

	err = userCenterServices.Auth(postForm.UserName, postForm.PassWord)
	if err != nil {
		utils.JsonFailedResponse(c, stateCode.UsernamePasswordUnMatch, nil)
		return
	}

	user, err := userServices.GetUserByStudentID(postForm.UserName)
	if err != nil && user != nil {
		user.OpenID = session.OpenID
		user.JHPassword = postForm.PassWord
	} else {
		user = &models.User{OpenID: session.OpenID, JHPassword: postForm.PassWord, StudentID: postForm.UserName}
		database.DB.Create(&user)
	}

	utils.JsonSuccessResponse(c, nil)
}
func BindZFPassword(c *gin.Context) {

}
func BindLibraryPassword(c *gin.Context) {

}

func BindSchoolCardPassword(c *gin.Context) {

}

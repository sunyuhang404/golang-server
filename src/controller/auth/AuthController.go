package auth

import (
	"bs-server/src/bo"
	"bs-server/src/model"
	"bs-server/src/response"
	"bs-server/src/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 登录
func Login(context *gin.Context) {
	var loginBO bo.LoginBO

	//
	if err := context.BindJSON(&loginBO); err != nil {
		context.JSON(http.StatusOK, response.ErrorParams.WithMsg(err.Error()))
		return
	}

	valid := validation.Validation{}

	//a := model.Auth{
	//	Username: loginBO.Username,
	//	Password: loginBO.Password,
	//}

	ok, _ := valid.Valid(&loginBO)

	data := make(map[string]interface{})

	if ok {
		isExist := model.CheckAuth(loginBO.Username, loginBO.Password)

		if isExist {
			token, err := util.GetToken(loginBO.Username, loginBO.Password)

			if err != nil {
				context.JSON(http.StatusOK, response.Error.WithData(err))
			} else {
				data["token"] = token
				context.JSON(http.StatusOK, response.OK.WithData(data))
			}
		}
	} else {
		var errList []interface{} = make([]interface{}, 0)

		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			errList = append(errList, err.Message)
		}

		context.JSON(http.StatusOK, response.Error.WithData(errList))
	}
}

// 注册
func Register(context *gin.Context) {}

// 获取用户信息
func GetUserInfo(context *gin.Context) {
	// context.Request.Header.Get("username")
}

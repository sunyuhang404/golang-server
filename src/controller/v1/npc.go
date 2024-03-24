package v1

import (
	"bs-server/src/model"
	"bs-server/src/response"
	"bs-server/src/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetNpcList(context *gin.Context) {

	params := make(map[string]interface{})

	data := make(map[string]interface{})

	//var state int = -1
	//if arg := context.Query("state"); arg != "" {
	//	state = com.StrTo(arg).MustInt()
	//	params["state"] = state
	//}

	data["list"] = model.GetNpcList(util.GetPage(context), util.PageSize, params)
	data["total"] = model.GetNpcTotal(params)

	context.JSON(http.StatusOK, response.OK.WithData(data))
}

func GetNpcInfo(context *gin.Context) {}

func CreateNpc(context *gin.Context) {

}

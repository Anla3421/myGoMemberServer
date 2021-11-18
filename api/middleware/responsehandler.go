package middleware

import (
	"net/http"
	"server/api/protocol"
	"server/env"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ResponseHanlder :
// 1. 程式開頭需先將 Response 寫入 gin.Context 中
//   c.Set(env.APIResKeyInGinContext, res)
// 2. 程式發生錯誤時, 應將錯誤寫入 gin.error 中, 並直接回傳
//   c.Error(err)
//   return
func ResponseHanlder() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Process request,進入點
		c.Next()
		// 取得API回應
		// Controller 執行時, 需先設定 c.Set(env.APIResKeyInGinContext, res)
		res := &protocol.Response{}

		// 設定一個空的物件 ， 用於有error時回傳一個空的result
		type initResult struct{}
		result := new(initResult)

		if infRes, ok := c.Get(env.APIResKeyInGinContext); ok {
			if realRes, isProtocol := infRes.(*protocol.Response); isProtocol {
				res = realRes
			}
		}
		// 取得執行過程中發生的錯誤, 只處理第一個
		lastError := c.Errors.Last()
		if lastError == nil {
			switch s := c.Writer.Status(); s {

			case http.StatusNotFound: //404
				// 當RESTFul API的路由找不到時, 會進入這個案例
				res.Code = http.StatusNotFound
				res.Message = "404 ,page not exists!"
				res.Result = result
				c.JSON(s, res)
			default: //200
				// 成功執行完畢, 回傳成功訊息
				if c.Request.URL.Path == "/swagger/doc.json" {
					return
				}

				res.Code = 1
				if res.Message == "" {
					res.Message = "Success"
				}
				if res.Result == nil {
					res.Result = result
				}
				c.JSON(s, res)
			}

			return
		}
		err := lastError.Err

		// 使用validator來處理Controller在c.ShouldBind(req),產生的Error
		if vErrors, ok := err.(validator.ValidationErrors); ok {
			res.Code = 400
			res.Message = "Wrong input, check again"
			res.Result = vErrors.Error()
			c.JSON(http.StatusBadRequest, res)
			return
		}

		// 自定義的Error類型
		res.Code = 500000
		res.Message = err.Error()
		res.Result = result
		c.JSON(http.StatusOK, res)
		return
	}
}

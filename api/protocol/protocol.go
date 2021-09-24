// Package protocol : api相關協定，包含可以回傳的error code訊息.
package protocol

// Response : api 回應格式
type Response struct {
	Code    int         `json:"Code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

// Code generated by goctl. DO NOT EDIT.
package types

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Page    interface{} `json:"page,omitempty"`
}

type InfoReq struct {
	Id string `json:"id"`
}

type InfoReply struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

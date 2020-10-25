// Code generated by goctl. DO NOT EDIT!
// Source: library.proto

package library

import "errors"

var errJsonConvert = errors.New("json convert error")

type (
	FindBookReply struct {
		No          string `json:"no,omitempty"`
		Name        string `json:"name,omitempty"`
		Author      string `json:"author,omitempty"`
		PublishFate string `json:"publishFate,omitempty"`
	}

	FindBookReq struct {
		// 书籍名称
		Name string `json:"name,omitempty"`
	}
)

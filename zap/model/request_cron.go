package model

import (
	"github.com/HXSecurity/Dongtai_USB/service"
)

type Response struct {
	Err  interface{} `json:"err"`
	Msg  string      `json:"msg"`
	Data []struct {
		VulName         string                    `json:"vul_name"`
		Detail          string                    `json:"detail"`
		VulLevel        string                    `json:"vul_level"`
		Urls            []string                  `json:"urls"`
		Payload         string                    `json:"payload"`
		CreateTime      int64                     `json:"create_time"`
		VulType         int                       `json:"vul_type"`
		RequestMessages []service.RequestMessages `json:"request_messages"`
	} `json:"data"`
}

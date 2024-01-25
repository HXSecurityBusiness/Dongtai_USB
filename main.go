package main

import (
	"github.com/HXSecurity/Dongtai_USB/config"
	xray "github.com/HXSecurity/Dongtai_USB/xray/request"
	zap "github.com/HXSecurity/Dongtai_USB/zap/request"
)

var usb = new(config.USB_config)
var USB_Xray = new(xray.USB_Xray)
var USB_Zap = new(zap.USB_Zap)

func main() {
	USB := usb.Init()
	router := USB.Group("api").Use(usb.JWTAuth())

	//推流模式(webhook)：
	router.POST("/v1/xray", USB_Xray.Xray)
	//拉流模式(cron):
	usb.Cron("xray", USB_Xray.Xray_cron)
	usb.Cron("zap", USB_Zap.Zap_cron)

	config.Log.Printf("The USB runs on port 5005")
	USB.Run(":5005")
}

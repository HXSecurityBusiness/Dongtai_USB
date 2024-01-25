package request

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/HXSecurity/Dongtai_USB/config"
	"github.com/HXSecurity/Dongtai_USB/service"
	"github.com/HXSecurity/Dongtai_USB/zap/model"
)

func (s *USB_Zap) Zap_cron(before time.Time, after time.Time) {
	var zapUrl = config.Viper.GetString("usb.zap_url")
	if zapUrl == "" {
		return
	}
	config.Log.Printf("正在自动从 zap 拉取数据 !!!")

	var zapResponse model.Response

	req, err := http.NewRequest("GET", zapUrl, nil)
	if err != nil {
		config.Log.Printf("http.Get: %v\n", err)
		return
	}

	// req.Header.Set("token", config.Viper.GetString("usb.xray_token"))
	req.Header.Set("Accept", "application/json, text/plain, */*")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		config.Log.Printf("client: %v\n", err)
		return
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		config.Log.Println(err)
		return
	}
	json.Unmarshal(body, &zapResponse)

	for _, data := range zapResponse.Data {
		res, err := engine_Zap.ReadHTTP(data.RequestMessages)
		if err != nil {
			config.Log.Print(err)
			return
		}

		agent := res[0].Response.Header.Get("Dt-Request-Id")
		if agent == "" {
			config.Log.Printf("找不到 Dt-Request-Id 请求头")
			return
		}

		engine := engine_Zap.EngineZap(agent, res)

		response := &service.Response{
			VulName:         data.Urls[0] + " " + data.VulName,
			Detail:          data.Detail,
			VulLevel:        data.VulLevel,
			Urls:            engine.Urls,
			Payload:         data.Payload,
			CreateTime:      data.CreateTime,
			VulType:         data.VulName,
			RequestMessages: data.RequestMessages,
			Target:          data.Urls[0],
			DtUUIDID:        engine.DtuuidID,
			AgentID:         engine.AgentID,
			DongtaiVulType:  []string{model.GetVulType(data.VulType, data.VulName)},
			Dtmark:          engine.Dtmark,
			DastTag:         "ZAP",
		}

		resResponse, err := json.Marshal(response)
		if err != nil {
			config.Log.Printf("无法解析json")
			return
		}
		config.Log.Print(string(resResponse))
		config.Log.Print(service.Client(response))
	}
}

package engine

import (
	"bufio"
	"io"
	"net/http"
	"strings"

	"github.com/HXSecurity/Dongtai_USB/config"
	"github.com/HXSecurity/Dongtai_USB/service"
	"github.com/HXSecurity/Dongtai_USB/zap/model"
)

type Engine_Zap struct {
}

func (engine *Engine_Zap) EngineZap(agent string, connections []model.Connection) model.Engine {
	var zap model.Engine
	AgentID := make([]string, 0)
	DtuuidID := make([]string, 0)
	Dtmark := make([]string, 0)
	for _, connection := range connections {
		dtmark := connection.Request.Header.Get("dt-mark-header")
		url := connection.Request.URL.String()
		if strings.Contains(url, "?") {
			URL_arr := strings.Split(url, "?")
			zap.Urls = append(zap.Urls, URL_arr[0])
		} else {
			zap.Urls = append(zap.Urls, url)
		}
		if agent == "" {
			config.Log.Printf("找不到 Dt-Request-Id 请求头")
			zap.AgentID = AgentID
			zap.DtuuidID = DtuuidID
		} else {
			arr := strings.Split(agent, ".")
			zap.AgentID = append(zap.AgentID, arr[0])
			zap.DtuuidID = append(zap.DtuuidID, arr[1])
		}
		if dtmark == "" {
			config.Log.Printf("找不到 dt-mark-header 请求头")
			zap.Dtmark = Dtmark
		} else {
			zap.Dtmark = append(zap.Dtmark, dtmark)
		}

	}
	return zap
}

func (engine *Engine_Zap) ReadHTTP(rms []service.RequestMessages) ([]model.Connection, error) {
	stream := make([]model.Connection, 0)
	for _, rm := range rms {
		req, err := http.ReadRequest(bufio.NewReader(strings.NewReader(rm.Request)))
		if err == io.EOF {
			print("error")
		}
		res, err := http.ReadResponse(bufio.NewReader(strings.NewReader(rm.Response)), req)
		if err == io.EOF {
			print("error")
		}
		stream = append(stream, model.Connection{Request: req, Response: res})
	}
	return stream, nil
}

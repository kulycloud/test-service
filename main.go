package main

import (
	"encoding/json"
	commonHttp "github.com/kulycloud/common/http"
	"github.com/kulycloud/common/logging"
	protoHttp "github.com/kulycloud/protocol/http"
)

var logger = logging.GetForComponent("service")

func main() {
	srv := commonHttp.NewServer(30006, testHandler)

	err := srv.Serve()

	if err != nil {
		logger.Panicw("could not serve", "error", err)
	}
}

type testResponseType struct {
	IncomingBody string                            `json:"incomingBody"`
	Method       string                            `json:"method"`
	Path         string                            `json:"path"`
	Headers      commonHttp.Headers                `json:"headers"`
	Source       string                            `json:"source"`
	KulyData     *protoHttp.RequestHeader_KulyData `json:"kulyData"`
	ServiceData  map[string]string                 `json:"serviceData"`
}

func testHandler(request *commonHttp.Request) *commonHttp.Response {
	if request.Path == "/echo" {
		return echoHandler(request)
	} else {
		return rootHandler(request)
	}
}

func echoHandler(request *commonHttp.Request) *commonHttp.Response {
	res := commonHttp.NewResponse()
	res.Headers.Set("Content-Type", request.Headers.Get("Content-Type"))
	res.Body = request.Body

	return res
}

func rootHandler(request *commonHttp.Request) *commonHttp.Response {
	body := request.Body.ReadAll()

	resData := testResponseType{
		IncomingBody: body.String(),
		Method:       request.Method,
		Path:         request.Path,
		Headers:      request.Headers,
		Source:       request.Source,
		KulyData:     request.KulyData,
		ServiceData:  request.ServiceData,
	}

	bodyJson, err := json.Marshal(resData)

	if err != nil {
		respErr := commonHttp.NewResponse()
		respErr.Status = 500
		return respErr
	}

	resp := commonHttp.NewResponse()
	resp.Headers.Set("X-MyHeader", "I set this! :)")
	resp.Headers.Set("Content-Type", "application/json")
	resp.Status = 200
	resp.Body = commonHttp.NewBody()
	resp.Body.Write(bodyJson)

	return resp
}

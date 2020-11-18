package main

import (
	"encoding/json"
	commonHttp "github.com/kulycloud/common/http"
	"github.com/kulycloud/common/logging"
	protoHttp "github.com/kulycloud/protocol/http"
)

var logger = logging.GetForComponent("service")

func main() {
	srv := commonHttp.NewHttpServer(30006, testHandler)

	err := srv.Serve()

	if err != nil {
		logger.Panicw("could not serve", "error", err)
	}
}

type testResponseType struct {
	IncomingBody string                            `json:"incomingBody"`
	HttpData     *protoHttp.RequestHeader_HttpData `json:"httpData"`
	KulyData     *protoHttp.RequestHeader_KulyData `json:"kulyData"`
	ServiceData  map[string]string                 `json:"serviceData"`
}

func testHandler(request *commonHttp.HttpRequest) *commonHttp.HttpResponse {
	if request.HttpData.Path == "/echo" {
		return echoHandler(request)
	} else {
		return rootHandler(request)
	}
}

func echoHandler(request *commonHttp.HttpRequest) *commonHttp.HttpResponse {
	res := commonHttp.NewHttpResponse()
	res.Headers["Content-Type"] = request.HttpData.Headers["Content-Type"]
	res.Body = request.Body

	return res
}

func rootHandler(request *commonHttp.HttpRequest) *commonHttp.HttpResponse {
	body := request.Body.ReadAll()

	resData := testResponseType {
		IncomingBody: body.String(),
		HttpData:     request.HttpData,
		KulyData:     request.KulyData,
		ServiceData:  request.ServiceData,
	}

	bodyJson, err := json.Marshal(resData)

	if err != nil {
		respErr := commonHttp.NewHttpResponse()
		respErr.Status = 500
		return respErr
	}

	resp := commonHttp.NewHttpResponse()
	resp.Headers["X-MyHeader"] = "I set this! :)"
	resp.Headers["Content-Type"] = "application/json"
	resp.Status = 200
	resp.Body = commonHttp.NewBody()
	resp.Body.Write(bodyJson)

	return resp
}

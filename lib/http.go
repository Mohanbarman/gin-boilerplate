package lib

import "github.com/gin-gonic/gin"

type H = map[string]interface{}

// Http response builder
type HttpResponseStruct struct {
	status int
	body   H
}

func (h *HttpResponseStruct) Data(data H) *HttpResponseStruct {
	h.body["data"] = data
	return h
}

func (h *HttpResponseStruct) Errors(e H) *HttpResponseStruct {
	h.body["errors"] = e
	return h
}

func (h *HttpResponseStruct) Message(m string) *HttpResponseStruct {
	h.body["message"] = m
	return h
}

func (h *HttpResponseStruct) Send(c *gin.Context) {
	c.JSON(h.status, h.body)
}

func HttpResponse(statusCode int) *HttpResponseStruct {
	return &HttpResponseStruct{
		status: statusCode,
		body: H{
			"code":    statusCode,
			"success": statusCode <= 400,
		},
	}
}

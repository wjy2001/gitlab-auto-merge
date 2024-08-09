package httpP

import (
	"github.com/go-resty/resty/v2"
	"log"
	"sync/atomic"
)

var Sum int64

type PreRequest struct {
	client *resty.Client
	Option RequestOption
}

type RequestOption struct {
	Method      string
	Url         string
	Headers     map[string]string
	QueryParams map[string]string
	Body        any
}

type InitRequest struct {
	BaseURL         string
	BaseHeaders     map[string]string
	BaseQueryParams map[string]string
}

// NewPreRequestClient 初始化链接
func NewPreRequestClient(initRequest InitRequest) *resty.Client {
	return resty.New().SetBaseURL(initRequest.BaseURL).SetHeaders(initRequest.BaseHeaders).SetQueryParams(initRequest.BaseQueryParams)
}

// NewPreRequest 初始化请求
func NewPreRequest(cli *resty.Client, opt RequestOption) *PreRequest {
	return &PreRequest{
		client: cli,
		Option: opt,
	}
}

func (r *PreRequest) newReq() *resty.Request {
	return r.client.R().SetHeaders(r.Option.Headers).SetQueryParams(r.Option.QueryParams).SetBody(r.Option.Body)
}

func (r *PreRequest) getResp(req *resty.Request) (*resty.Response, error) {
	return req.Execute(r.Option.Method, r.Option.Url)
}

func (r *PreRequest) GetRespBody() ([]byte, error) {
	sum := atomic.AddInt64(&Sum, 1)
	if sum%10 == 0 {
		log.Println("sum:", sum)
	}
	req := r.newReq()
	resp, err := r.getResp(req)
	return resp.Body(), err
}

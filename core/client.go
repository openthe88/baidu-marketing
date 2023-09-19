package core

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/openthe88/baidu-marketing/core/internal/debug"
	"github.com/openthe88/baidu-marketing/model"
)

// SDKClient  object
type SDKClient struct {
	token     string
	ocpcToken string
	username  string
	password  string
	debug     bool
	header    map[string]string // Custom header map.
}

// NewSDKClient init sdk client
func NewSDKClient(token string, ocpcToken string) *SDKClient {
	return &SDKClient{
		token:     token,
		ocpcToken: ocpcToken,
		header:    make(map[string]string),
	}
}

// Token get token
func (c *SDKClient) Token() string {
	return c.token
}

// OcpcToken get ocpc token
func (c *SDKClient) OcpcToken() string {
	return c.ocpcToken
}

// SetUser set username password
func (c *SDKClient) SetUser(username string, password string) {
	c.username = username
	c.password = password
}

// SetDebug set debug mode
func (c *SDKClient) SetDebug(debug bool) {
	c.debug = debug
}

// SetHeader sets a custom HTTP header pair for the client.
func (c *SDKClient) SetHeader(key, value string) {
	c.header[key] = value
}

func (c *SDKClient) UpFile(url string, resp interface{}, fileReq *model.FileRequest, req map[string]string) error {
	// 创建一个字节缓冲区来构建请求体
	requestBody := &bytes.Buffer{}
	writer := multipart.NewWriter(requestBody)

	part, err := writer.CreateFormFile(fileReq.Filed, fileReq.Name)
	if err != nil {
		return err
	}

	_, err = io.Copy(part, fileReq.File)
	if err != nil {
		return err
	}

	for k, v := range req {
		writer.WriteField(k, v)
	}

	writer.Close()

	// 创建HTTP POST请求
	request, err := http.NewRequest("POST", url, requestBody)
	if err != nil {
		return err
	}

	// 设置请求头，指定Content-Type为multipart/form-data
	request.Header.Set("Content-Type", writer.FormDataContentType())

	// 发送请求
	client := &http.Client{
		Timeout: time.Second * 200, // 设置请求超时时间
	}
	httpResp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	err = debug.DecodeJSONHttpResponse(httpResp.Body, resp, c.debug)
	if err != nil {
		debug.PrintError(err, c.debug)
		return err
	}

	return nil
}

func (c *SDKClient) DoFile(req *model.Request, resp interface{}, fileReq *model.FileRequest, reqBody map[string]string) (*model.ResponseHeader, error) {
	var reqResp model.Response
	err := c.UpFile(req.Url(), &reqResp, fileReq, reqBody)
	if err != nil {
		return nil, err
	}
	if reqResp.IsError() {
		if reqResp.Header.Status == 1 && resp != nil {
			err = json.Unmarshal(reqResp.Body, resp)
			if err != nil {
				return &reqResp.Header, err
			}
		}
		return &reqResp.Header, reqResp
	}
	if resp != nil {
		err = json.Unmarshal(reqResp.Body, resp)
		if err != nil {
			return &reqResp.Header, err
		}
	}
	return &reqResp.Header, nil
}

// Do execute api request
func (c *SDKClient) Do(req *model.Request, resp interface{}) (*model.ResponseHeader, error) {
	if req.Header.Token == "" {
		req.Header.Token = c.token
	}
	if req.Header.AccessToken == "" {
		if req.Header.Username == "" {
			req.Header.Username = c.username
		}
		if req.Header.Password == "" {
			req.Header.Password = c.password
		}
	}
	reqBytes, _ := json.Marshal(req)
	var reqResp model.Response
	err := c.Post(req.Url(), reqBytes, &reqResp)
	if err != nil {
		return nil, err
	}
	if reqResp.IsError() {
		if reqResp.Header.Status == 1 && resp != nil {
			err = json.Unmarshal(reqResp.Body, resp)
			if err != nil {
				return &reqResp.Header, err
			}
		}
		return &reqResp.Header, reqResp
	}
	if resp != nil {
		err = json.Unmarshal(reqResp.Body, resp)
		if err != nil {
			return &reqResp.Header, err
		}
	}
	return &reqResp.Header, nil
}

// Post data through api
func (c *SDKClient) Post(reqUrl string, reqBytes []byte, resp interface{}) error {
	debug.PrintPostJSONRequest(reqUrl, reqBytes, c.debug)
	httpReq, err := http.NewRequest("POST", reqUrl, bytes.NewReader(reqBytes))
	if err != nil {
		return err
	}
	contType, ok := c.header["Content-Type"]
	if !ok {
		httpReq.Header.Add("Content-Type", "application/json;charset=utf-8")
	} else {
		httpReq.Header.Add("Content-Type", contType)
	}

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	err = debug.DecodeJSONHttpResponse(httpResp.Body, resp, c.debug)
	if err != nil {
		debug.PrintError(err, c.debug)
		return err
	}
	return nil
}

package core

import (
	"bytes"
	"encoding/json"
	"net/http"

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
	}
}

// Token get token
func (c SDKClient) Token() string {
	return c.token
}

// OcpcToken get ocpc token
func (c SDKClient) OcpcToken() string {
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

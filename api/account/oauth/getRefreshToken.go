package oauth

import (
	"encoding/json"
	"github.com/openthe88/baidu-marketing/core"
	"github.com/openthe88/baidu-marketing/model/account/oauth"
)

func GetRefreshToken(clt *core.SDKClient, req oauth.GetRefreshTokenRequest) (oauth.GetRefreshTokenResponse, error) {
	reqBytes, _ := json.Marshal(req)
	url := req.Url()

	var resp oauth.GetRefreshTokenResponse
	err := clt.Post(url, reqBytes, &resp)
	return resp, err
}

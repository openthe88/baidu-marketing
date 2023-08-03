package oauth

import (
	"encoding/json"
	"errors"
	"github.com/openthe88/baidu-marketing/core"
	"github.com/openthe88/baidu-marketing/model/account/oauth"
)

func GetRefreshToken(clt *core.SDKClient, req oauth.GetRefreshTokenRequest) (oauth.RefreshTokenData, error) {
	reqBytes, _ := json.Marshal(req)
	url := req.Url()

	var resp oauth.GetRefreshTokenResponse
	err := clt.Post(url, reqBytes, &resp)

	if resp.Code != 0 {
		return oauth.RefreshTokenData{}, errors.New(resp.Message)
	}

	return resp.Data, err
}

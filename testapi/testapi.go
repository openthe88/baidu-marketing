package main

import (
	"fmt"
	"github.com/openthe88/baidu-marketing/api/account/oauth"
	"github.com/openthe88/baidu-marketing/core"
	OauthMod "github.com/openthe88/baidu-marketing/model/account/oauth"
)

var (
	Appid           = `433fe9a4071ffb303035c9fa08160503`
	SecretKey       = `9081f94869fc5f6f84ffe6929d279832`
	UserId    int64 = 49053238
	Scope           = `1007690`
	AuthCode        = `eyJhbGciOiJIUzM4NCJ9.eyJhdWQiOiLmipXmlL7ns7vnu58iLCJzdWIiOiJleGMiLCJ1aWQiOjQ5MDUzMjM4LCJhcHBJZCI6IjQzM2ZlOWE0MDcxZmZiMzAzMDM1YzlmYTA4MTYwNTAzIiwiaXNzIjoi5ZWG5Lia5byA5Y-R6ICF5Lit5b-DIiwicGxhdGZvcm1JZCI6IjQ5NjAzNDU5NjU5NTg1NjE3OTQiLCJleHAiOjE2OTEwMzM2NzIsImp0aSI6Ijc3NzE5NzkwMjk5MjUwMDMyOTMifQ.zzB8zCpJZ8guLbrksGZ25rixCAa1DdYLaOO-v_oaSnSyNi_0wHfHuA5mE_tkN_lI`
	Token           = `eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJhY2MiLCJhdWQiOiLmipXmlL7ns7vnu58iLCJ1aWQiOjQ5MDUzMjM4LCJhcHBJZCI6IjQzM2ZlOWE0MDcxZmZiMzAzMDM1YzlmYTA4MTYwNTAzIiwiaXNzIjoi5ZWG5Lia5byA5Y-R6ICF5Lit5b-DIiwicGxhdGZvcm1JZCI6IjQ5NjAzNDU5NjU5NTg1jE3OTQiLCJleHAiOjE2OTExMTYyMjUsImp0aSI6Ijc3NzE5OTE4MTE3NDc3NzQ0NjkifQ.cBm_FGwDWreACVfzdX5YiUlgwm0-j7OxF9Z6DB94PquCXCPFeQbIpY-yni8_VPdO`
	ReToken         = `eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJyZWYiLCJhdWQiOiLokavoiqblqIMiLCJ1aWQiOjQ5MDQwODkxLCJhcHBJZCI6Ijc2MDZiMWUwZDdhMmViYzA0ZWVmZDJhYTQ1MjJmZTEwIiwiaXNzIjoi5ZWG5Lia5byA5Y-R6ICF5Lit5b-DIiwia2V5SWQiOjYwMjUxNSwicGxhdGZvcm1JZCI6IjQ5NjAzNDU5NjU5NTg1NjE3OTQiLCJleHAiOjE2OTM2MzMyNDYsImp0aSI6Ijc3NzE5ODY0NTE2Mjg0OTg5NTAifQ.r_ftI1mAj_jj9VuMpXi6rfXwiCvUK3Qdh1sv-A87n3qN5M-9tBAi6KWzKQE315Ld`
	OpenId          = `7771957314570403840`
)

func main() {
	//oauthUrl()
	//getToken()
	//getReToken()
}

func oauthUrl() {
	req := OauthMod.GetAppOauthRequest{
		AppId:    Appid,
		Scope:    Scope,
		State:    "lhd_test",
		Callback: "http://market-api.gzfei.com/api/v1/advertiser/baidu_oauth/callback",
	}
	OauthUrl, _ := oauth.GetAppOauth(req)

	fmt.Println(OauthUrl)
}

func getToken() {

	sdk := &core.SDKClient{}
	sdk.SetDebug(true)
	req := OauthMod.GetAccessTokenRequest{
		AppId:     Appid,
		AuthCode:  AuthCode,
		SecretKey: SecretKey,
		GrantType: "auth_code",
		UserId:    UserId,
	}

	token, err := oauth.GetAccessToken(sdk, req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("token为%s\n", token.Data.AccessToken)
	fmt.Printf("刷新token为%s\n", token.Data.RefreshToken)
	fmt.Printf("openid为%s\n", token.Data.OpenId)
	fmt.Printf("userid为%d\n", token.Data.UserId)

}

func getReToken() {
	sdk := &core.SDKClient{}
	sdk.SetDebug(true)
	req := OauthMod.GetRefreshTokenRequest{
		AppId:        Appid,
		RefreshToken: ReToken,
		SecretKey:    SecretKey,
		UserId:       UserId,
	}

	token, err := oauth.GetRefreshToken(sdk, req)
	if err != nil {
		return
	}

	fmt.Printf("token为%s\n", token.Data.AccessToken)
	fmt.Printf("reToken为%s\n", token.Data.RefreshToken)
}

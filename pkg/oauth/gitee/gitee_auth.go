package gitee

import (
	"context"
	"demo01/util"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/goburrow/cache"

	"golang.org/x/oauth2"
)

var ctxCache = cache.New(cache.WithMaximumSize(1000), cache.WithExpireAfterAccess(10*time.Minute))

type UserInfo struct {
	Id        int    `json:"id"`
	Login     string `json:"login"`
	NodeId    string `json:"nodeId"`
	AvatarUrl string `json:"avatar_url"`
	Url       string `json:"url"`
	HtmlUrl   string `json:"html_url"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Bio       string `json:"bio"`
	Company   string `json:"company"`
	Blog      string `json:"blog"`
	Location  string `json:"location"`
}

// params callback携带的参数
func newOauthConfig(redirectUrl string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     "0e7d12d5dbb2d666c6e179301ff73ea2722fd29f01289f5c7699da21765ebe27",
		ClientSecret: "59ddf446365fa48817eb35f5789090ffe28f4206f763bf8abaa3f3161adb7bb8",
		RedirectURL:  redirectUrl,
		Scopes:       []string{},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://gitee.com/oauth/authorize",
			TokenURL: "https://gitee.com/oauth/token",
		},
	}
}

func AuthCodeURL(params map[string]string) string {
	// 将跳转地址写入上线文
	state := util.GenUUID("")
	redirectUrl := getRedirectUrl(params)
	ctxCache.Put(state, redirectUrl)

	return newOauthConfig(redirectUrl).AuthCodeURL(state)
}

// 获取回调跳转地址
func getRedirectUrl(params map[string]string) string {
	redirectUrl := "http://localhost:9123/api/v1/oauth/gitee/callback"
	if len(params) > 0 {
		ub := util.ParseUrl(redirectUrl)
		ub.AddQueries(params)
		redirectUrl = ub.Build()
	}
	return redirectUrl
}

// 根据code获取用户信息
// 流程为先使用code换取accessToken，然后根据accessToken获取用户信息
func GetUserInfoByCode(code, state string) (*UserInfo, error) {
	// 从上下文中获取跳转地址
	val, found := ctxCache.GetIfPresent(state)
	var redirectUrl string
	if found {
		redirectUrl = val.(string)
	}

	token, err := newOauthConfig(redirectUrl).Exchange(context.TODO(), code)
	if err != nil {
		return nil, err
	}
	return GetUserInfo(token.AccessToken)
}

// 根据accessToken获取用户信息
func GetUserInfo(accessToken string) (*UserInfo, error) {
	response, err := resty.New().R().SetQueryParam("access_token", accessToken).Get("https://gitee.com/api/v5/user")
	if err != nil {
		fmt.Printf("Get user info error %s", err.Error())
		return nil, err
	}
	content := string(response.Body())
	fmt.Printf("Gitee userInfo: %v", content)

	userInfo := &UserInfo{}
	err = json.Unmarshal([]byte(content), userInfo)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

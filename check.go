package go_pay

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type checker struct {
	googlePackageName  string
	googleClientId     string
	googleClientSecret string
	googleRedirectUri  string
	googleRefreshToken string
	tokenData          *accessToken
}

type accessToken struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	CreatedAt    int64  `json:"created_at"`
}

type Options func(c *checker)

func WithPackageName(pn string) Options {
	return func(c *checker) {
		c.googlePackageName = pn
	}
}

func WithClientId(clientId string) Options {
	return func(c *checker) {
		c.googleClientId = clientId
	}
}

func WithRedirectUri(redirectUri string) Options {
	return func(c *checker) {
		c.googleRedirectUri = redirectUri
	}
}

func WithClientSecret(clientSecret string) Options {
	return func(c *checker) {
		c.googleClientSecret = clientSecret
	}
}

func WithGoogleAuthCode(code string) Options {
	return func(c *checker) {
		_, err := c.getGoogleRefreshToken(code)
		if err != nil {
			panic(err)
		}
	}
}

func NewChecker(options ...Options) *checker {
	c := &checker{}
	for _, o := range options {
		o(c)
	}
	return c
}

//从苹果服务器校验, receipt为Base64编码格式
func (c *checker) CheckAppleReceipt(receipt string) (*ApplePayResponse, error) {
	var (
		err          error
		data         []byte
		res          *ApplePayResponse
		request      *http.Request
		checkUrl     = productionVerifyUrl
		receiptParam = `{"receipt-data": "` + receipt + `"}`
		checkCnt     = 1
	)

label:
	request, err = http.NewRequest("POST", checkUrl, strings.NewReader(receiptParam))
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status code: %d", response.StatusCode)
	}
	defer response.Body.Close()
	if data, err = ioutil.ReadAll(response.Body); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	switch res.Status {
	case 0:
		return res, nil
	case 21007: //需要转入沙盒测试环境
		checkUrl = sandboxVerifyUrl
		goto label
	case 21002, 21005, 21009: //需要重试
		if checkCnt == 3 {
			return nil, fmt.Errorf("apple check time out, quit with status code: %d", res.Status)
		}
		checkCnt++
		goto label
	default: //校验失败
		return nil, fmt.Errorf("apple check fail with status code: %d", res.Status)
	}
}

func (c *checker) CheckGooglePayToken(token, productId string) (*GooglePayResponse, error) {
	//判断access_token是否过期
	if c.tokenData == nil {
		return nil, ErrInvalidToken
	}

	if time.Now().Unix()-c.tokenData.CreatedAt >= c.tokenData.ExpiresIn {
		if err := c.refreshGoogleAccessToken(); err != nil {
			return nil, err
		}
	}

	queryUrl := fmt.Sprintf("https://www.googleapis.com/androidpublisher/v3/applications/%s/purchases/products/%s/tokens/%s?access_token=%s", c.googlePackageName, productId, token, c.tokenData.AccessToken)
	response, err := http.Get(queryUrl)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status code: %d", response.StatusCode)
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result *GooglePayResponse
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *checker) refreshGoogleAccessToken() error {
	param := fmt.Sprintf("grant_type=refresh_token&client_id=%s&client_secret=%s&refresh_token=%s", c.googleClientId, c.googleClientSecret, c.tokenData.RefreshToken)
	response, err := http.Post(googleBaseUrl, "application/x-www-form-urlencoded", strings.NewReader(param))
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("http status code: %d", response.StatusCode)
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(data, &c.tokenData); err != nil {
		return err
	}
	c.tokenData.CreatedAt = time.Now().Unix()
	return nil
}

func (c *checker) getGoogleRefreshToken(code string) (string, error) {
	param := fmt.Sprintf("client_id=%s&client_secret=%s&redirect_uri=%s&grant_type=authorization_code&code=%s", c.googleClientId, c.googleClientSecret, c.googleRedirectUri, code)

	response, err := http.Post(googleBaseUrl, "application/x-www-form-urlencoded", strings.NewReader(param))
	if err != nil {
		return "", err
	}
	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("http status code: %d", response.StatusCode)
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	if err = json.Unmarshal(data, &c.tokenData); err != nil {
		return "", err
	}
	c.tokenData.CreatedAt = time.Now().Unix()
	return c.tokenData.RefreshToken, nil
}

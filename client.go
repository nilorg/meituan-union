package union

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Client 客户端
type Client struct {
	conf       *Config
	httpClient *http.Client
}

// NewClient ...
func NewClient(conf *Config) (client *Client, err error) {
	client = &Client{
		conf:       conf,
		httpClient: &http.Client{},
	}
	return
}

// Config ...
func (c *Client) Config() Config {
	return *c.conf
}

// ResponseStatus 响应状态
type ResponseStatus struct {
	Status int    `json:"status"` // 状态值，0为成功，非0为异常
	Des    string `json:"des"`    // 异常描述信息
}

// getExecute 执行
func (c *Client) getExecute(uri string, param interface{}) (body []byte, err error) {
	var values url.Values
	values, err = SignStructToParameter(param)
	if err != nil {
		return
	}
	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s?%s", c.conf.BaseURL, uri, values.Encode()), nil)
	if err != nil {
		return
	}
	var resp *http.Response
	resp, err = c.httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err = fmt.Errorf("请求接口状态码错误：%d", resp.StatusCode)
		return
	}
	body, err = ioutil.ReadAll(resp.Body)
	return
}

// postExecute 执行
func (c *Client) postExecute(uri string, param interface{}) (body []byte, err error) {
	var values url.Values
	values, err = SignStructToParameter(param)
	if err != nil {
		return
	}
	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, c.conf.BaseURL+uri, strings.NewReader(values.Encode()))
	if err != nil {
		return
	}
	var resp *http.Response
	resp, err = c.httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err = fmt.Errorf("请求接口状态码错误：%d", resp.StatusCode)
		return
	}
	body, err = ioutil.ReadAll(resp.Body)
	return
}

// GenerateLink 自助取链接口（新版）
func (c *Client) GenerateLink(req *GenerateLinkRequest) (resp *GenerateLinkResponse, err error) {
	req.AppKey = c.conf.AppKey
	err = req.SignMD5(c.conf.SignatureKey)
	if err != nil {
		return
	}
	var body []byte
	body, err = c.getExecute("/generateLink", req)
	if err != nil {
		return
	}
	resp = new(GenerateLinkResponse)
	err = json.Unmarshal(body, resp)
	if err != nil {
		resp = nil
		return
	}
	if resp.Status != 0 {
		err = fmt.Errorf("请求接口错误：%s", resp.Des)
		return
	}
	return
}

// OrderList 自助取链接口（新版）
func (c *Client) OrderList(req *OrderListRequest) (resp *OrderListResponse, err error) {
	req.AppKey = c.conf.AppKey
	err = req.SignMD5(c.conf.SignatureKey)
	if err != nil {
		return
	}
	var body []byte
	body, err = c.getExecute("/orderList", req)
	if err != nil {
		return
	}
	resp = new(OrderListResponse)
	err = json.Unmarshal(body, resp)
	if err != nil {
		resp = nil
		return
	}
	return
}

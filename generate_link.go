package union

// https://union.meituan.com/v2/apiDetail?id=25

// GenerateLinkRequest 自助取链接口（新版）
type GenerateLinkRequest struct {
	AppKey    string `json:"appkey"`    // 媒体
	ActID     string `json:"actId"`     // 活动id
	SID       string `json:"sid"`       // 推广位sid，支持通过接口自定义创建
	LinkType  string `json:"linkType"`  // 链接类型，枚举值：1.h5链接	2.deeplink(唤起)链接 3.中间页唤起链接 4.微信小程序唤起路径
	ShortLink string `json:"shortLink"` // 0表示获取长链	1表示获取短链
	Sign      string `json:"sign"`      // 请求签名
}

// SignMD5 md5
func (req *GenerateLinkRequest) SignMD5(signatureKey string) error {
	params, err := SignStructToParameter(*req)
	if err != nil {
		return err
	}
	value := SignMD5(params, signatureKey)
	req.Sign = value
	return nil
}

// GenerateLinkResponse ...
type GenerateLinkResponse struct {
	ResponseStatus
	Data string `json:"data"`
}

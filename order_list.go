package union

// https://union.meituan.com/v2/apiDetail?id=23

// OrderListRequest 订单列表查询接口（新版）
type OrderListRequest struct {
	AppKey        string `json:"appkey"`        // 媒体
	Ts            string `json:"ts"`            // 请求时刻10位时间戳(秒级)，有效期60s
	ActID         string `json:"actId"`         // 活动id
	BusinessLine  string `json:"businessLine"`  // 业务线
	StartTime     string `json:"startTime"`     // 查询起始时间10位时间戳，以下单时间为准
	EndTime       string `json:"endTime"`       // 查询截止时间10位时间戳，以下单时间为准
	Page          string `json:"page"`          // 分页参数，起始值从1开始
	Limit         string `json:"limit"`         // 每页显示数据条数，最大值为100
	QueryTimeType string `json:"queryTimeType"` // 查询时间类型，枚举值 1 按订单支付时间查询
	Sign          string `json:"sign"`          // 请求签名
}

// SignMD5 md5
func (req *OrderListRequest) SignMD5(signatureKey string) error {
	params, err := SignStructToParameter(*req)
	if err != nil {
		return err
	}
	value := SignMD5(params, signatureKey)
	req.Sign = value
	return nil
}

// OrderListResponse ...
type OrderListResponse struct {
	DataList []*OrderDataList `json:"dataList"`
	Total    int              `json:"total"`
	Msg      string           `json:"msg"`
}

type OrderDataList struct {
	OrderID                     string `json:"orderid"`
	PayTime                     string `json:"paytime"`
	PayPrice                    string `json:"payprice"`
	SID                         string `json:"sid"`
	SMSTitle                    string `json:"smstitle"`
	AppKey                      string `json:"appkey"`
	Status                      int    `json:"status"`
	Profit                      string `json:"profit"`
	CpaProfit                   string `json:"cpaProfit"`
	RefundTime                  string `json:"refundtime"`
	RefundPrice                 string `json:"refundprice"`
	RefundProfit                string `json:"refundprofit"`
	CpaRefundProfit             string `json:"cpaRefundProfit"`
	Extra                       string `json:"extra"`
	TradeTypeList               []int  `json:"tradeTypeList"`
	TradeTypeBusinessTypeMapStr string `json:"tradeTypeBusinessTypeMapStr"`
	RiskOrder                   int    `json:"riskOrder"`
	BusinessLine                int    `json:"businessLine"`
	SubBusinessLine             int    `json:"subBusinessLine"`
	ActID                       int    `json:"actId"`
}

package models

import "time"

type AuthResult struct {
	Result  string `json:"result"`
	Msg     string `json:"msg"`
	RetCode int    `json:"ret_code"`
}

type UnbindResult struct {
	Result string `json:"result"`
	Msg    string `json:"msg"`
}

type CheckerResult struct {
	Result string `json:"result"`
	Time   string `json:"time"`
	Msg    string `json:"msg"`
}

type NetInfo struct {
	UserName      string `json:"user_name"`     // 姓名
	UsedTime      string `json:"used_time"`     // 上网时间
	UsedFlow      string `json:"used_flow"`     // 使用流量
	Balance       string `json:"balance"`       // 余额
	Account       string `json:"account"`       // 学号
	Status        string `json:"status"`        // 状态
	Package       string `json:"package"`       // 用户类型
	PackageDetail string `json:"packageDetail"` // 用户类型详情
	BillingType   string `json:"billingType"`   // 付费模式
	BillingCycle  string `json:"billingCycle"`  // 计费周期
}

type LoginHistoryRecord struct {
	LoginTime    time.Time `json:"loginTime"`    // 上线时间
	LogoutTime   time.Time `json:"logoutTime"`   // 注销时间
	IP           string    `json:"ip"`           // IP地址
	MAC          string    `json:"mac"`          // MAC信息（大写无分隔符）
	UseTime      int       `json:"useTime"`      // 使用时长（分钟）
	UseFlow      float64   `json:"useFlow"`      // 使用流量（M）
	BillingType  string    `json:"billingType"`  // 计费方式（1时长 2流量 3包月）
	ChargeAmount float64   `json:"chargeAmount"` // 计费金额
	HostName     string    `json:"hostName"`     // 主机名
	TerminalType string    `json:"terminalType"` // 终端类型（如PC、移动终端）
}

type OnlineRecord struct {
	Brasid       string    `json:"brasid"`       // BRAS ID
	DownFlow     int64     `json:"downFlow"`     // 下行流量（单位：KB？根据原始响应值推测）
	HostName     string    `json:"hostName"`     // 主机名
	IP           string    `json:"ip"`           // IP地址
	LoginTime    time.Time `json:"loginTime"`    // 上线时间
	MAC          string    `json:"mac"`          // MAC地址（无分隔符）
	SessionId    string    `json:"sessionId"`    // 会话ID
	TerminalType string    `json:"terminalType"` // 终端类型（可能带 "#" 前缀）
	UpFlow       int64     `json:"upFlow"`       // 上行流量
	UseTime      int       `json:"useTime"`      // 使用时长（单位：秒）
	UserId       int64     `json:"userId"`       // 用户ID
}

type LoginData struct {
	StuID    string
	Password string
	UA       string
	Isp      string
	IPAddr   string
	MACAddr  string
}

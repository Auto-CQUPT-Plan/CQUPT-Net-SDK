package spyder

import (
	"crypto/tls"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/Auto-CQUPT-Plan/CQUPT-Net-SDK/core/models"
	"github.com/bytedance/sonic"
	"github.com/gocolly/colly"
)

func newUnsafeTransport() *http.Transport {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // 跳过证书验证
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	return transport
}

func generateUrlWithQueryParamsForOnlineList() (string, error) {
	t := rand.Float64()
	_timestamp := time.Now().UnixMilli()

	u := fmt.Sprintf("%s?t=%f&order=asc&_=%d", onlineListUrl, t, _timestamp)
	return u, nil
}

func generateRequestForOnlineList(data *models.LoginData, url string, cookie string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Cookie", cookie)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("User-Agent", uaMap[data.UA])
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func generateUrlWithQueryParamsForLoginHistory() (*url.URL, error) {
	u, err := url.Parse(loginHistoryUrl)
	if err != nil {
		return nil, fmt.Errorf("parse url error: %v", err.Error())
	}

	return u, nil
}

func generateRequestForLoginHistory(u *url.URL, data *models.LoginData, cookie string) (*http.Request, error) {
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Cookie", cookie)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("User-Agent", uaMap[data.UA])
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")

	return req, nil
}

func generateUrlWithQueryParamsForEnactiveCookie() (*url.URL, error) {
	u, err := url.Parse(enactiveCookieUrl)
	if err != nil {
		return nil, fmt.Errorf("parse url error: %v", err.Error())
	}

	return u, nil
}

func generateRequestForEnactiveCookie(u *url.URL, data *models.LoginData, cookie string) (*http.Request, error) {
	if _, ok := uaMap[data.UA]; !ok {
		return nil, fmt.Errorf("unsupported ua: %v", uaMap[data.UA])
	}

	request, _ := http.NewRequest(http.MethodPost, u.String(), nil)
	request.Header.Add("Host", serviceHost)
	request.Header.Add("Referer", serviceRefer)
	request.Header.Add("Cookie", cookie)
	request.Header.Add("User-Agent", uaMap[data.UA])
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("X-Requested-With", "XMLHttpRequest")

	request.Header.Add("Origin", "https://202.202.32.120:8443")
	request.Header.Add("Cache-Control", "max-age=0")
	request.Header.Add("Sec-Fetch-Site", "same-origin")
	request.Header.Add("Sec-Fetch-Mode", "no-cors")
	request.Header.Add("Sec-Fetch-User", "?1")
	request.Header.Add("Sec-Fetch-Dest", "image")

	return request, nil
}

func generateUrlWithQueryParamsForServiceLogin() (*url.URL, error) {
	u, err := url.Parse(serviceUrl)
	if err != nil {
		return nil, fmt.Errorf("parse url error: %v", err.Error())
	}

	return u, nil
}

func generateRequestForServiceLogin(u *url.URL, data *models.LoginData, cookie string, checkCode string) (*http.Request, error) {
	if _, ok := uaMap[data.UA]; !ok {
		return nil, fmt.Errorf("unsupported ua: %v", uaMap[data.UA])
	}

	formData := url.Values{}
	formData.Set("foo", "")
	formData.Set("bar", "")
	formData.Set("checkcode", checkCode)
	formData.Set("account", data.StuID)
	formData.Set("password", data.Password)
	formData.Set("code", "")

	request, _ := http.NewRequest(http.MethodPost, u.String(), strings.NewReader(formData.Encode()))
	request.Header.Add("Host", serviceHost)
	request.Header.Add("Referer", serviceRefer)
	request.Header.Add("Cookie", cookie)
	request.Header.Add("User-Agent", uaMap[data.UA])
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("X-Requested-With", "XMLHttpRequest")

	request.Header.Add("Origin", "https://202.202.32.120:8443")
	request.Header.Add("Cache-Control", "max-age=0")
	request.Header.Add("Sec-Fetch-Site", "same-origin")
	request.Header.Add("Sec-Fetch-Mode", "navigate")
	request.Header.Add("Sec-Fetch-User", "?1")
	request.Header.Add("Sec-Fetch-Dest", "document")

	return request, nil
}

func (r *Spyder) LoginCquptService(data *models.LoginData) (string, error) {
	var cookie string
	var checkCode string

	// 获取新Cookie
	collector := colly.NewCollector(colly.UserAgent(uaMap[data.UA]), colly.IgnoreRobotsTxt())
	collector.WithTransport(newUnsafeTransport())
	collector.OnResponse(func(r *colly.Response) {
		cookies := r.Headers.Values("Set-Cookie")
		fmt.Printf("Set-Cookie: %v \n", cookies)
		if len(cookies) > 0 {
			parts := strings.Split(cookies[0], ";")
			cookie = strings.TrimSpace(parts[0])
		}
	})
	collector.OnHTML("input[type='hidden'][name='checkcode']", func(e *colly.HTMLElement) { checkCode = e.Attr("value") })
	err := collector.Visit(loginUrl)
	if err != nil {
		return "", err
	}
	collector.Wait()

	fmt.Printf("cookie: %s, code: %s \n", cookie, checkCode)

	// 激活Cookie
	u, err := generateUrlWithQueryParamsForEnactiveCookie()
	if err != nil {
		return "", err
	}

	request, err := generateRequestForEnactiveCookie(u, data, cookie)
	if err != nil {
		return "", err
	}

	response, err := r.HttpRequester.Do(request)
	if err != nil {
		return "", fmt.Errorf("do request error: %s", err.Error())
	}
	if response.Body != nil {
		response.Body.Close()
	}

	// 执行登录
	u, err = generateUrlWithQueryParamsForServiceLogin()
	if err != nil {
		return "", err
	}

	fmt.Printf("url: %s \n", u.String())

	request, err = generateRequestForServiceLogin(u, data, cookie, checkCode)
	if err != nil {
		return "", err
	}

	fmt.Printf("requestBody: %s ,cookies: %v \n", request.Body, request.Cookies())

	response, err = r.HttpRequester.Do(request)
	if err != nil {
		return "", fmt.Errorf("do request error: %s", err.Error())
	}

	if lo := response.Header.Get("Location"); lo != "/Self/dashboard" {
		return "", fmt.Errorf("login failed, location: %s", lo)
	} else {
		fmt.Printf("success location: %s \n", lo)
	}

	return cookie, nil
}

func (r *Spyder) GetNetServiceInfo(data *models.LoginData) (*models.NetInfo, error) {
	var cookie string
	var info models.NetInfo

	// 清理所有空白字符（空格、换行、制表符等）
	cleanAllWhitespace := func(s string) string {
		return strings.Map(func(r rune) rune {
			if unicode.IsSpace(r) {
				return -1 // 丢弃空白字符
			}
			return r
		}, s)
	}

	if ck := r.Cache.GetItem("Cookie"); ck == nil {
		newCookie, err := r.LoginCquptService(data)
		if err != nil {
			return nil, err
		}
		r.Cache.AddItem("Cookie", newCookie, 15*time.Second)
		cookie = newCookie
	} else {
		cookie = ck.(string)
	}

	fmt.Printf("GetNetServiceInfo-cookie: %s \n", cookie)

	collector := colly.NewCollector(colly.UserAgent(uaMap[data.UA]), colly.IgnoreRobotsTxt())
	collector.WithTransport(newUnsafeTransport())
	collector.OnRequest(func(r *colly.Request) { r.Headers.Set("Cookie", cookie) })

	collector.OnHTML("div.thumbnail div.caption", func(e *colly.HTMLElement) {
		// 提取用户名和问候
		h4Text := e.ChildText("h4")
		parts := strings.Split(h4Text, "，") // 注意这里是中文逗号
		if len(parts) >= 2 {
			info.UserName = cleanAllWhitespace(parts[0])
		}
	})

	// 提取统计数据（已用时长、流量、余额）
	collector.OnHTML("div.user-info1 dl", func(e *colly.HTMLElement) {
		dtText := e.ChildText("dt")
		ddText := e.ChildText("dd")

		switch cleanAllWhitespace(ddText) {
		case "已用时长":
			info.UsedTime = cleanAllWhitespace(dtText)
		case "已用流量":
			info.UsedFlow = cleanAllWhitespace(dtText)
		case "账户余额":
			info.Balance = cleanAllWhitespace(dtText)
		}
	})

	// 提取详细信息（账号、状态、套餐等）
	collector.OnHTML("div.panel-body div.row", func(e *colly.HTMLElement) {
		label := e.ChildText("label")
		value := cleanAllWhitespace(e.ChildText("span"))

		switch {
		case strings.Contains(label, "账　　号"):
			info.Account = value
		case strings.Contains(label, "状　　态"):
			info.Status = cleanAllWhitespace(e.ChildText("span.label"))
		case strings.Contains(label, "套　　餐"):
			info.Package = cleanAllWhitespace(e.ChildText("span"))
			info.PackageDetail = cleanAllWhitespace(e.ChildText("small"))
		case strings.Contains(label, "计费方式"):
			info.BillingType = value
		case strings.Contains(label, "计费周期"):
			var dates []string
			e.ForEach("span.label", func(_ int, el *colly.HTMLElement) {
				dates = append(dates, cleanAllWhitespace(el.Text))
			})
			if len(dates) >= 2 {
				info.BillingCycle = dates[0] + " 至 " + dates[1]
			}
		}
	})

	err := collector.Visit(getNetInfoUrl)
	if err != nil {
		return nil, err
	}
	collector.Wait()

	return &info, nil
}

func (r *Spyder) GetLoginHistory(data *models.LoginData) ([]*models.LoginHistoryRecord, error) {
	var cookie string

	if ck := r.Cache.GetItem("Cookie"); ck == nil {
		newCookie, err := r.LoginCquptService(data)
		if err != nil {
			return nil, err
		}
		r.Cache.AddItem("Cookie", newCookie, 15*time.Second)
		cookie = newCookie
	} else {
		cookie = ck.(string)
	}

	u, err := generateUrlWithQueryParamsForLoginHistory()
	if err != nil {
		return nil, err
	}

	req, err := generateRequestForLoginHistory(u, data, cookie)
	if err != nil {
		return nil, err
	}

	resp, err := r.HttpRequester.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rawRecords [][]interface{}
	if err := sonic.Unmarshal(body, &rawRecords); err != nil {
		return nil, err
	}

	records := make([]*models.LoginHistoryRecord, 0, len(rawRecords))
	for _, raw := range rawRecords {
		if len(raw) < 12 {
			continue
		}
		record := &models.LoginHistoryRecord{}

		if loginTimeMs, ok := raw[0].(float64); ok {
			record.LoginTime = time.UnixMilli(int64(loginTimeMs))
		}

		if logoutTimeMs, ok := raw[1].(float64); ok {
			record.LogoutTime = time.UnixMilli(int64(logoutTimeMs))
		}

		if ip, ok := raw[2].(string); ok {
			record.IP = ip
		}

		if mac, ok := raw[3].(string); ok {
			record.MAC = mac
		}

		if useTime, ok := raw[4].(float64); ok {
			record.UseTime = int(useTime)
		}

		if useFlow, ok := raw[5].(float64); ok {
			record.UseFlow = useFlow
		}

		if billingTypeNum, ok := raw[6].(float64); ok {
			switch int(billingTypeNum) {
			case 1:
				record.BillingType = "时长"
			case 2:
				record.BillingType = "流量"
			case 3:
				record.BillingType = "包月"
			default:
				record.BillingType = "未知"
			}
		}

		if chargeAmount, ok := raw[7].(float64); ok {
			record.ChargeAmount = chargeAmount
		}

		if raw[8] != nil {
			if hostName, ok := raw[8].(string); ok {
				record.HostName = hostName
			}
		}

		if terminalType, ok := raw[10].(string); ok {
			record.TerminalType = terminalType
		}

		records = append(records, record)
	}

	return records, nil
}

func (r *Spyder) GetOnlineList(data *models.LoginData) ([]*models.OnlineRecord, error) {
	var cookie string

	// 从缓存获取Cookie，若不存在则登录
	if ck := r.Cache.GetItem("Cookie"); ck == nil {
		newCookie, err := r.LoginCquptService(data)
		if err != nil {
			return nil, err
		}
		r.Cache.AddItem("Cookie", newCookie, 15*time.Second)
		cookie = newCookie
	} else {
		cookie = ck.(string)
	}

	// 构造URL（带随机参数和时间戳）
	u, err := generateUrlWithQueryParamsForOnlineList()
	if err != nil {
		return nil, err
	}

	// 生成请求
	req, err := generateRequestForOnlineList(data, u, cookie)
	if err != nil {
		return nil, err
	}

	// 发送请求
	resp, err := r.HttpRequester.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 解析JSON数组
	var rawRecords []map[string]interface{}
	if err := sonic.Unmarshal(body, &rawRecords); err != nil {
		return nil, err
	}

	records := make([]*models.OnlineRecord, 0, len(rawRecords))
	for _, raw := range rawRecords {
		record := &models.OnlineRecord{}

		// 逐个字段解析
		if v, ok := raw["brasid"].(string); ok {
			record.Brasid = v
		}
		if v, ok := raw["downFlow"].(string); ok {
			if val, err := strconv.ParseInt(v, 10, 64); err == nil {
				record.DownFlow = val
			}
		}
		if v, ok := raw["hostName"].(string); ok {
			record.HostName = v
		}
		if v, ok := raw["ip"].(string); ok {
			record.IP = v
		}
		if v, ok := raw["loginTime"].(string); ok {
			// 解析时间格式 "2026-03-28 20:56:38"
			t, err := time.ParseInLocation("2006-01-02 15:04:05", v, time.Local)
			if err == nil {
				record.LoginTime = t
			}
		}
		if v, ok := raw["mac"].(string); ok {
			record.MAC = v
		}
		if v, ok := raw["sessionId"].(string); ok {
			record.SessionId = v
		}
		if v, ok := raw["terminalType"].(string); ok {
			record.TerminalType = v
		}
		if v, ok := raw["upFlow"].(string); ok {
			if val, err := strconv.ParseInt(v, 10, 64); err == nil {
				record.UpFlow = val
			}
		}
		if v, ok := raw["useTime"].(string); ok {
			if val, err := strconv.Atoi(v); err == nil {
				record.UseTime = val
			}
		}
		if v, ok := raw["userId"].(float64); ok {
			record.UserId = int64(v)
		}

		records = append(records, record)
	}

	return records, nil
}

package spyder

const (
	host              = "192.168.200.2:801"
	serviceHost       = "202.202.32.120:8443"
	referer           = "http://192.168.200.2/"
	serviceRefer      = "https://202.202.32.120:8443/Self/login/?302=LI"
	requestUrl        = "http://192.168.200.2:801/eportal/"
	loginUrl          = "https://202.202.32.120:8443/Self/login/?302=LI"
	serviceUrl        = "https://202.202.32.120:8443/Self/login/verify"
	getNetInfoUrl     = "https://202.202.32.120:8443/Self/dashboard"
	enactiveCookieUrl = "https://202.202.32.120:8443/Self/login/randomCode?t=0.5562441101050084"
	loginHistoryUrl   = "https://202.202.32.120:8443/Self/dashboard/getLoginHistory"
	onlineListUrl     = "https://202.202.32.120:8443/Self/dashboard/getOnlineList"
	c                 = "Portal"
	a                 = "login"
	aUnbind           = "unbind_mac"
	aChecker          = "page_type_data"
	callback          = "dr1003"
	callbackUnbind    = "dr1002"
	callbackChecker   = "dr1001"
	loginMethod       = "1"
	jsVersion         = "3.3.3"
)

var (
	uaMap = map[string]string{
		"phone":   "Mozilla/5.0 (Linux; Android 12; Pixel 6 Build/SD1A.210817.023; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/94.0.4606.71 Mobile Safari/537.36",
		"pad":     "Mozilla/5.0 (iPad; CPU OS 17_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Safari/605.1.15 Chrome/117.0.5938.62",
		"desktop": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.5938.62 Safari/537.36",
	}

	ispMap = map[string]struct{}{
		"telecom": {},
		"cmcc":    {},
		"unicom":  {},
		"xyw":     {},
	}
)

package spyder

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"

	"github.com/Auto-CQUPT-Plan/CQUPT-Net-SDK/core/models"
	"github.com/bytedance/sonic"
)

func generateUrlWithQueryParamsForLogin(data *models.LoginData) (*url.URL, error) {
	var device uint8 = 1

	u, err := url.Parse(requestUrl)
	if err != nil {
		return nil, fmt.Errorf("parse url error: %v", err.Error())
	}

	if data.StuID == "" {
		return nil, fmt.Errorf("empty stuID")
	}
	if data.Password == "" {
		return nil, fmt.Errorf("empty password")
	}
	if data.IPAddr == "" {
		return nil, fmt.Errorf("empty ipAddr")
	}
	if data.MACAddr == "" {
		return nil, fmt.Errorf("empty macAddr")
	}
	if isp, ok := ispMap[data.Isp]; !ok {
		return nil, fmt.Errorf("unsupported isp: %v", isp)
	}

	query := u.Query()
	query.Add("c", c)
	query.Add("a", a)
	query.Add("callback", callback)
	query.Add("login_method", loginMethod)

	if data.UA == "desktop" {
		device = 0
	}
	query.Add("user_account", fmt.Sprintf(",%d,%s@%s", device, data.StuID, data.Isp))
	query.Add("user_password", data.Password)
	query.Add("wlan_user_ip", data.IPAddr)
	query.Add("wlan_user_mac", strings.ReplaceAll(data.MACAddr, ":", ""))
	query.Add("jsVersion", jsVersion)
	u.RawQuery = query.Encode()

	return u, nil
}

func generateRequestForLogin(u *url.URL, data *models.LoginData) (*http.Request, error) {
	if ua, ok := uaMap[data.UA]; !ok {
		return nil, fmt.Errorf("unsupported ua: %v", ua)
	}

	request, _ := http.NewRequest(http.MethodGet, u.String(), nil)
	request.Header.Add("Host", host)
	request.Header.Add("Referer", referer)
	request.Header.Add("User-Agent", data.UA)

	return request, nil
}

func generateUrlWithQueryParamsForLogout(data *models.LoginData) (*url.URL, error) {
	u, err := url.Parse(requestUrl)
	if err != nil {
		return nil, fmt.Errorf("parse url error: %v", err.Error())
	}

	if data.StuID == "" {
		return nil, fmt.Errorf("empty stuID")
	}
	if data.Isp == "" {
		return nil, fmt.Errorf("empty isp")
	}
	if data.IPAddr == "" {
		return nil, fmt.Errorf("empty ipAddr")
	}
	if data.MACAddr == "" {
		return nil, fmt.Errorf("empty macAddr")
	}
	if _, ok := ispMap[data.Isp]; !ok {
		return nil, fmt.Errorf("unsupported isp: %v", data.Isp)
	}

	query := u.Query()

	query.Add("c", c)
	query.Add("a", aUnbind)
	query.Add("callback", callbackUnbind)
	query.Add("user_account", fmt.Sprintf("%s@%s", data.StuID, data.Isp))
	query.Add("wlan_user_mac", strings.ReplaceAll(data.MACAddr, ":", ""))
	query.Add("wlan_user_ip", data.IPAddr)
	query.Add("jsVersion", jsVersion)
	query.Add("v", fmt.Sprintf("%d", rand.Intn(9000)+1000))

	u.RawQuery = query.Encode()

	return u, nil
}

func generateRequestForLogout(u *url.URL, data *models.LoginData) (*http.Request, error) {
	if ua, ok := uaMap[data.UA]; !ok {
		return nil, fmt.Errorf("unsupported ua: %v", ua)
	}

	request, _ := http.NewRequest(http.MethodGet, u.String(), nil)
	request.Header.Add("Host", host)
	request.Header.Add("Referer", referer)
	request.Header.Add("User-Agent", data.UA)

	return request, nil
}

func generateUrlWithQueryParamsForChecker() (*url.URL, error) {
	u, err := url.Parse(requestUrl)
	if err != nil {
		return nil, fmt.Errorf("parse url error: %v", err.Error())
	}

	query := u.Query()
	query.Add("c", c)
	query.Add("a", aChecker)
	query.Add("callback", callbackChecker)
	query.Add("v", fmt.Sprintf("%d", rand.Intn(9000)+1000))

	u.RawQuery = query.Encode()

	return u, nil
}

func generateRequestForChecker(u *url.URL, data *models.LoginData) (*http.Request, error) {
	if ua, ok := uaMap[data.UA]; !ok {
		return nil, fmt.Errorf("unsupported ua: %v", ua)
	}

	request, _ := http.NewRequest(http.MethodGet, u.String(), nil)
	request.Header.Add("Host", host)
	request.Header.Add("Referer", referer)
	request.Header.Add("User-Agent", data.UA)

	return request, nil
}

func (r *Spyder) NetLogin(data *models.LoginData) (*models.AuthResult, error) {
	var result models.AuthResult

	u, err := generateUrlWithQueryParamsForLogin(data)
	if err != nil {
		return nil, err
	}

	request, err := generateRequestForLogin(u, data)
	if err != nil {
		return nil, err
	}

	response, err := r.HttpRequester.Do(request)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(response.Body)
	if err != nil {
		return nil, fmt.Errorf("do request error: %s", err.Error())
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("read raw response error: %s", err.Error())
	}

	resultString := string(bytes)
	err = sonic.Unmarshal([]byte(resultString[7:len(resultString)-1]), &result)
	if err != nil {
		return nil, fmt.Errorf("nnmarshal data error: %s", err.Error())
	}

	if result.Result == "0" && result.RetCode == 2 {
		return &models.AuthResult{Result: result.Result, Msg: "当前设备已认证", RetCode: result.RetCode}, nil
	}

	return &result, nil
}

func (r *Spyder) NetLogout(data *models.LoginData) (*models.UnbindResult, error) {
	var result models.UnbindResult

	u, err := generateUrlWithQueryParamsForLogout(data)
	fmt.Println(u.String())
	if err != nil {
		return nil, err
	}

	request, err := generateRequestForLogout(u, data)
	if err != nil {
		return nil, err
	}

	response, err := r.HttpRequester.Do(request)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(response.Body)
	if err != nil {
		return nil, fmt.Errorf("do request error: %s", err.Error())
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("read raw response error: %s", err.Error())
	}

	resultString := string(bytes)
	err = sonic.Unmarshal([]byte(resultString[7:len(resultString)-1]), &result)
	if err != nil {
		return nil, fmt.Errorf("nnmarshal data error: %s", err.Error())
	}

	return &result, nil
}

func (r *Spyder) NetChecker(data *models.LoginData) (*models.CheckerResult, error) {
	var result models.CheckerResult

	u, err := generateUrlWithQueryParamsForChecker()
	fmt.Println(u.String())
	if err != nil {
		return nil, err
	}

	request, err := generateRequestForChecker(u, data)
	if err != nil {
		return nil, err
	}

	response, err := r.HttpRequester.Do(request)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(response.Body)
	if err != nil {
		return nil, fmt.Errorf("do request error: %s", err.Error())
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("read raw response error: %s", err.Error())
	}

	resultString := string(bytes)
	err = sonic.Unmarshal([]byte(resultString[7:len(resultString)-1]), &result)
	if err != nil {
		return nil, fmt.Errorf("nnmarshal data error: %s", err.Error())
	}

	return &result, nil
}

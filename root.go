package CQUPT_Net_SDK

import (
	"crypto/tls"
	"net/http"

	"github.com/Auto-CQUPT-Plan/CQUPT-Net-SDK/core/spyder"
	"github.com/Auto-CQUPT-Plan/CQUPT-Net-SDK/core/utils"
)

type SDK struct {
	spyder *spyder.Spyder
}

func NewSDK() *SDK {
	// http
	client := &http.Client{
		Transport:     &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},               // 取消TLS
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse }, // 取消自动重定向
	}
	// cache
	cache := utils.NewTimerDB()

	// core
	s := &spyder.Spyder{HttpRequester: client, Cache: cache}

	return &SDK{spyder: s}
}

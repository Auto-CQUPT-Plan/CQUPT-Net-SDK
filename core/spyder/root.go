package spyder

import (
	"net/http"

	"github.com/Auto-CQUPT-Plan/CQUPT-Net-SDK/core/utils"
)

type Spyder struct {
	HttpRequester *http.Client
	Cache         *utils.TimerDB
}
